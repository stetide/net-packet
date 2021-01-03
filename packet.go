package main

/*
#include <sys/socket.h>
#include <arpa/inet.h>
#include <net/ethernet.h>
#include <net/if.h>
#include <linux/if_packet.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/time.h>
#include <stdio.h>
#include <errno.h>
#include <sys/signal.h>

#define E_TIMEOUT 11
#define E_INTERRUPT 4

typedef unsigned char byte;
typedef byte ip4[4];
typedef byte mac[6];

struct arp_header {
    unsigned short int hwaddr_type;
    unsigned short int proto_type;
    byte hwaddr_len;
    byte proto_len;
    unsigned short int ar_op;
    mac src_mac_addr;
    ip4 src_ip_addr;
    mac tgt_mac_addr;
    ip4 tgt_ip_adddr;
};

// set_promisc sets promiscuous mode on sock
// according to the interface
// opt can be ether ADD_MEMBERSHIP or DROP_MEMBERSHIP
int set_promisc(int sock, int ifindx, int opt) {
    if (ifindx == 0) {
        return -1;
    }

    struct packet_mreq mreq;
    memset(&mreq, 0, sizeof(mreq));
    mreq.mr_ifindex = ifindx;
    mreq.mr_type = PACKET_MR_PROMISC;
    return setsockopt(sock, SOL_PACKET, opt, &mreq, sizeof(mreq));
}

// enable_promisc enables promiscuous mode on sock
int enable_promisc(int sock, int ifindx) {
    return set_promisc(sock, ifindx, PACKET_ADD_MEMBERSHIP);
}

// disable_promisc disables promiscuous mode on sock
int disable_promisc(int sock, int ifindx) {
    return set_promisc(sock, ifindx, PACKET_DROP_MEMBERSHIP);
}

int set_timeout(int sock, long microseconds) {
    struct timeval timeout;
    timeout.tv_sec = 0;
    timeout.tv_usec = microseconds;

    if (setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO, (char *)&timeout, sizeof(timeout)) == -1) {
        return -1;
    }
    if (setsockopt(sock, SOL_SOCKET, SO_SNDTIMEO, (char *)&timeout, sizeof(timeout)) == -1) {
        return -1;
    }
    return 0;
}

int set_send_timeout(int sock, long microseconds) {
    struct timeval timeout;
    timeout.tv_sec = 0;
    timeout.tv_usec = microseconds;

    return (setsockopt(sock, SOL_SOCKET, SO_SNDTIMEO, (char *)&timeout, sizeof(timeout)) == -1);
}

int set_recv_timeout(int sock, long microseconds) {
    struct timeval timeout;
    timeout.tv_sec = 0;
    timeout.tv_usec = microseconds;

    return (setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO, (char *)&timeout, sizeof(timeout)) == -1);
}

int create_socket(int protocol) {
	int sock = socket(AF_PACKET, SOCK_RAW, htons(protocol));
	return sock;
}

// close_socket disables promiscuous mode
// and closes socket
int close_socket(int sock) {
    return close(sock);
}

struct arp_header *arp_decode_packet(char *pkt) {
    struct ether_header *eth_hdr;
    struct arp_header *arp_hdr;

    eth_hdr = (struct ether_header *)pkt;
    arp_hdr = (struct arp_header *)(pkt + sizeof(struct ether_header));
    return arp_hdr;
}
*/
import "C"
import (
	"errors"
	"net"
	"syscall"
	"time"
)

// Socket handler
// export Socket
type Socket int

// NewSocket creates new packet socket
// export NewSocket
func NewSocket(protocol EtherProto) (Socket, error) {
	sock := int(C.create_socket(C.int(protocol)))
	if sock == -1 {
		return 0, errors.New("error on create socket")
	}
	return Socket(sock), nil
}

// SetPromisc enables or disables promiscuos mode on s
// export SetPromisc
func (s Socket) SetPromisc(iface net.Interface, promisc bool) error {
	if promisc {
		if int(C.enable_promisc(C.int(s), C.int(iface.Index))) == -1 {
			return errors.New("error on set_promisc")
		}
	} else {
		if int(C.disable_promisc(C.int(s), C.int(iface.Index))) == -1 {
			return errors.New("error on set_promisc")
		}
	}
	return nil
}

// SetTimeout sets a time out for send and recevie
// export SetTimeout
func (s Socket) SetTimeout(timeout time.Duration) error {
	usecs := timeout.Round(time.Microsecond).Microseconds()
	if int(C.set_timeout(C.int(s), C.long(usecs))) == -1 {
		return errors.New("error on set_timeout")
	}
	return nil
}

// SetSendTimeout sets a timeout on send
// export SetSendTimeout
func (s Socket) SetSendTimeout(timeout time.Duration) error {
	usecs := timeout.Round(time.Microsecond).Microseconds()
	if int(C.set_send_timeout(C.int(s), C.long(usecs))) == -1 {
		return errors.New("error on set_timeout")
	}
	return nil
}

// SetRecvTimeout sets a timeout on recv
// export SetRecvTimeout
func (s Socket) SetRecvTimeout(timeout time.Duration) error {
	usecs := timeout.Round(time.Microsecond).Microseconds()
	if int(C.set_recv_timeout(C.int(s), C.long(usecs))) == -1 {
		return errors.New("error on set_timeout")
	}
	return nil
}

// Close closes the socket
// export Close
func (s Socket) Close() error {
	if int(C.close_socket(C.int(s))) == -1 {
		return errors.New("error on close_socket")
	}
	return nil
}

// Read reads socket data into buf
// export Read
func (s Socket) Read(buf []byte) (int, error) {
	n, _, err := syscall.Recvfrom(int(s), buf, 0)
	return n, err
}

// Recv works like Read but lets you pass flags
// flags can be found in the syscall package
// export Recv
func (s Socket) Recv(buf []byte, flags int) (int, error) {
	n, _, err := syscall.Recvfrom(int(s), buf, flags)
	return n, err
}

func main() {}

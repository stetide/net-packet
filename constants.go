package main

// EtherProto represents Ethernet Protocols
// export EtherProto
type EtherProto int

const (
	/* ETH_P_LOOP Ethernet Loopback packet	*/
	ETH_P_LOOP EtherProto = 0x0060
	/* ETH_P_PUP Xerox PUP packet		*/
	ETH_P_PUP EtherProto = 0x0200
	/* Xerox PUP Addr Trans packet	*/
	ETH_P_PUPAT EtherProto = 0x0201
	/* TSN (IEEE 1722) packet	*/
	ETH_P_TSN EtherProto = 0x22F0
	/* ERSPAN version 2 (type III)	*/
	ETH_P_ERSPAN2 EtherProto = 0x22EB
	/* Internet Protocol packet	*/
	ETH_P_IP EtherProto = 0x0800
	/* CCITT X.25			*/
	ETH_P_X25 EtherProto = 0x0805
	/* Address Resolution packet	*/
	ETH_P_ARP EtherProto = 0x0806
	/* G8BPQ AX.25 Ethernet Packet	[ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_BPQ EtherProto = 0x08FF
	/* Xerox IEEE802.3 PUP packet */
	ETH_P_IEEEPUP EtherProto = 0x0a00
	/* Xerox IEEE802.3 PUP Addr Trans packet */
	ETH_P_IEEEPUPAT EtherProto = 0x0a01
	/* B.A.T.M.A.N.-Advanced packet [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_BATMAN EtherProto = 0x4305
	/* DEC Assigned proto           */
	ETH_P_DEC EtherProto = 0x6000
	/* DEC DNA Dump/Load            */
	ETH_P_DNA_DL EtherProto = 0x6001
	/* DEC DNA Remote Console       */
	ETH_P_DNA_RC EtherProto = 0x6002
	/* DEC DNA Routing              */
	ETH_P_DNA_RT EtherProto = 0x6003
	/* DEC LAT                      */
	ETH_P_LAT EtherProto = 0x6004
	/* DEC Diagnostics              */
	ETH_P_DIAG EtherProto = 0x6005
	/* DEC Customer use             */
	ETH_P_CUST EtherProto = 0x6006
	/* DEC Systems Comms Arch       */
	ETH_P_SCA EtherProto = 0x6007
	/* Trans Ether Bridging		*/
	ETH_P_TEB EtherProto = 0x6558
	/* Reverse Addr Res packet	*/
	ETH_P_RARP EtherProto = 0x8035
	/* Appletalk DDP		*/
	ETH_P_ATALK EtherProto = 0x809B
	/* Appletalk AARP		*/
	ETH_P_AARP EtherProto = 0x80F3
	/* 802.1Q VLAN Extended Header  */
	ETH_P_8021Q EtherProto = 0x8100
	/* ERSPAN type II		*/
	ETH_P_ERSPAN EtherProto = 0x88BE
	/* IPX over DIX			*/
	ETH_P_IPX EtherProto = 0x8137
	/* IPv6 over bluebook		*/
	ETH_P_IPV6 EtherProto = 0x86DD
	/* IEEE Pause frames. See 802.3 31B */
	ETH_P_PAUSE EtherProto = 0x8808
	/* Slow Protocol. See 802.3ad 43B */
	ETH_P_SLOW EtherProto = 0x8809
	/* Web-cache coordination protocol
	* defined in draft-wilson-wrec-wccp-v2-00.txt */
	ETH_P_WCCP EtherProto = 0x883E
	/* MPLS Unicast traffic		*/
	ETH_P_MPLS_UC EtherProto = 0x8847
	/* MPLS Multicast traffic	*/
	ETH_P_MPLS_MC EtherProto = 0x8848
	/* MultiProtocol Over ATM	*/
	ETH_P_ATMMPOA EtherProto = 0x884c
	/* PPPoE discovery messages     */
	ETH_P_PPP_DISC EtherProto = 0x8863
	/* PPPoE session messages	*/
	ETH_P_PPP_SES EtherProto = 0x8864
	/* HPNA, wlan link local tunnel */
	ETH_P_LINK_CTL EtherProto = 0x886c
	/* Frame-based ATM Transport
	* over Ethernet*/
	ETH_P_ATMFATE EtherProto = 0x8884
	/* Port Access Entity (IEEE 802.1X) */
	ETH_P_PAE EtherProto = 0x888E
	/* ATA over Ethernet		*/
	ETH_P_AOE EtherProto = 0x88A2
	/* 802.1ad Service VLAN		*/
	ETH_P_8021AD EtherProto = 0x88A8
	/* 802.1 Local Experimental 1.  */
	ETH_P_802_EX1 EtherProto = 0x88B5
	/* 802.11 Preauthentication */
	ETH_P_PREAUTH EtherProto = 0x88C7
	/* TIPC 			*/
	ETH_P_TIPC EtherProto = 0x88CA
	/* Link Layer Discovery Protocol */
	ETH_P_LLDP EtherProto = 0x88CC
	/* Media Redundancy Protocol	*/
	ETH_P_MRP EtherProto = 0x88E3
	/* 802.1ae MACsec */
	ETH_P_MACSEC EtherProto = 0x88E5
	/* 802.1ah Backbone Service Tag */
	ETH_P_8021AH EtherProto = 0x88E7
	/* 802.1Q MVRP                  */
	ETH_P_MVRP EtherProto = 0x88F5
	/* IEEE 1588 Timesync */
	ETH_P_1588 EtherProto = 0x88F7
	/* NCSI protocol		*/
	ETH_P_NCSI EtherProto = 0x88F8
	/* IEC 62439-3 PRP/HSRv0	*/
	ETH_P_PRP EtherProto = 0x88FB
	/* Fibre Channel over Ethernet  */
	ETH_P_FCOE EtherProto = 0x8906
	/* Infiniband over Ethernet	*/
	ETH_P_IBOE EtherProto = 0x8915
	/* TDLS */
	ETH_P_TDLS EtherProto = 0x890D
	/* FCoE Initialization Protocol */
	ETH_P_FIP EtherProto = 0x8914
	/* IEEE 802.21 Media Independent Handover Protocol */
	ETH_P_80221 EtherProto = 0x8917
	/* IEC 62439-3 HSRv1	*/
	ETH_P_HSR EtherProto = 0x892F
	/* Network Service Header */
	ETH_P_NSH EtherProto = 0x894F
	/* Ethernet loopback packet, per IEEE 802.3 */
	ETH_P_LOOPBACK EtherProto = 0x9000
	/* deprecated QinQ VLAN [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_QINQ1 EtherProto = 0x9100
	/* deprecated QinQ VLAN [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_QINQ2 EtherProto = 0x9200
	/* deprecated QinQ VLAN [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_QINQ3 EtherProto = 0x9300
	/* Ethertype DSA [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_EDSA EtherProto = 0xDADA
	/* Fake VLAN Header for DSA [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_DSA_8021Q EtherProto = 0xDADB
	/* ForCES inter-FE LFB type */
	ETH_P_IFE EtherProto = 0xED3E
	/* IBM af_iucv [ NOT AN OFFICIALLY REGISTERED ID ] */
	ETH_P_AF_IUCV EtherProto = 0xFBFB

	/* If the value in the ethernet type is less than this value
	* then the frame is Ethernet II. Else it is 802.3 */
	ETH_P_802_3_MIN EtherProto = 0x0600
)

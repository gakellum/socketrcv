# socketrcv
This sample Go app opens a socket to a sender and receives as many 1024-byte messages as it can.  It creates a go routine to handle each message.  Each go routine sleeps for 50ms to simulate processing time.  Its purpose is to test the data transfer rate via a socket to a receiving process on the same box.

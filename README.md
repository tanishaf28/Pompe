# Pompe: Byzantine Ordered Consensus without Byzantine Oligarchy 

This project implements a *Byzantine Fault Tolerant (BFT) ordered consensus mechanism* based on Pompe, ensuring that transaction ordering remains fair even in the presence of Byzantine nodes.  

🔹Overview:

Pompe is a Byzantine ordered consensus protocol that prevents a **Byzantine oligarchy**, ensuring that no small group of malicious nodes can control the transaction ordering. This implementation simulates the protocol with **honest and Byzantine replicas.  

🛠 How It Works
1. Replicas propose timestamps for a transaction.  
   - Honest replicas provide correct timestamps.  
   - Byzantine replicas can propose arbitrary (random) timestamps.  
2. Timestamps are sorted, ensuring a fair distribution.  
3. The median timestamp is selected, mitigating the influence of Byzantine nodes.  
4. Transactions are ordered** based on the median timestamp.  

Why Pompe?
Unlike traditional BFT protocols, **Pompe prevents Byzantine nodes from controlling transaction ordering**, ensuring fairness and decentralization in distributed systems.  

🔧 Run the Simulation  
Execute `main.go` to simulate transaction ordering for clients like Alice and Bob, demonstrating how Pompe ensures fair order in adversarial environments.  

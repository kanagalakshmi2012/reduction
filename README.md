# reduction
**EFFICIENT REDUCTION OF COMMIT DELAYS IN CONSENSUS BASED SYSTEMS**
* Author: Kanagalakshmi Murugan
* Published In : International Journal for Multidisciplinary Research (IJFMR)
* Publication Date: 06-2024
* E-ISSN: 2582-2160
* Impact Factor: 9.88
* Link:

**Abstract:**
This paper explores the impact of commit latency in etcd, a distributed key-value store that uses the Raft consensus algorithm and State Machine Replication (SMR) to ensure consistency. While SMR provides strong fault tolerance, it introduces significant latency as system size increases due to quorum-based communication. Commit latency, affected by factors like network delays and disk I/O, directly impacts system responsiveness. To address this, the study proposes optimizations using a Write-Ahead Log (WAL) mechanism to reduce delay in consensus without sacrificing consistency. The findings aim to enhance etcd’s performance in large-scale, latency-sensitive environments.

**Key Contributions:**
* Algorithm Development
  Optimized Write Ahead Log latency algorithm to achieve high performnace in etcd environment.
* Performance Comparison
  Conducted bench marking between State Machine Replication latencyc and Write Ahead Log latency. 
* Reserach Leadership
  Directed both the investigation and engineering efforts, emphasizing progress in distributed databases by introducing novel algorithmic approaches.

**Relevance & Real-World Impact**
* **Commit Latency Analysis:**\
Examined how commit latency in etcd, influenced by network delays and disk I/O, affected system responsiveness and overall performance as the cluster size grew.

* **WAL-Based Optimization:**\
Implemented Write-Ahead Log techniques to minimize consensus delays, improved throughput, and reduced latency while maintaining strong consistency in distributed key-value stores.
* **Academic Recognition :** \
    Included in academic papers and technical analyses focusing on DNS query handling and performance improvements in distributed key-value stores.
* **Educational Impact:** \
    Findings were incorporated into courses and research projects, supporting continuous academic dialogue on container orchestration and cloud system efficiency.

**Experimental Results (Summary)**:

  | Nodes | SMR Latency (ms) | WAL Latency (ms) | Improvement (%) |
  |-------|------------------| -----------------| ----------------|
  | 3     | 14.2             | 2.7              | 81.0            |
  | 5     | 18.4             | 3.0              | 83.7            |
  | 7     | 22.3             | 3.3              | 85.2            |
  | 9     | 26.1             | 3.6              | 86.2            |
  | 11    | 30.0             | 4.0              | 86.7            |

**Citation** \
EFFICIENT REDUCTION OF COMMIT DELAYS IN CONSENSUS BASED SYSTEMS \
Kanagalakshmi Murugan \
International Journal for Multidisciplinary Research \
E-ISSN-2582-2160 \
License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/kanagalakshmi-murugan-12b278199 | **Email**: kanagalakshmi2004@gmail.com

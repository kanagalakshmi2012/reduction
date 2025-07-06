# reduction
**EFFICIENT REDUCTION OF COMMIT DELAYS IN CONSENSUS BASED SYSTEMS**
* Author: Kanagalakshmi Murugan
* Published In : International Journal on Science and Technology (IJSAT)
* Publication Date: 06-2022
* E-ISSN: 2229-7677
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
* **Commit Latency Analysis:**
Examined how commit latency in etcd, influenced by network delays and disk I/O, affected system responsiveness and overall performance as the cluster size grew.

* **WAL-Based Optimization:**
Implemented Write-Ahead Log techniques to minimize consensus delays, improved throughput, and reduced latency while maintaining strong consistency in distributed key-value stores.
* **Academic Recognition :** \
    need to add here
* **Educational Impact:** \
    need to add here \

**Experimental Results (Summary)**:

  | Nodes | VR Replication Time (ms) | ZAB Replication Time (ms) | Improvement (%) |
  |-------|--------------------------| --------------------------| ----------------|
  | 3     | 6.0                      | 5.2                       | 13.3            |
  | 5     | 8.1                      | 6.7                       | 17.3            |
  | 7     | 10.5                     | 8.4                       | 20.0            |
  | 9     | 13.0                     | 10.3                      | 20.8            |
  | 11    | 15.6                     | 12.2                      | 21.8            |

**Citation** \
OPTIMIZING READ PERFORMANCE IN DISTRIBUTED systems USING Chrony Sync process. \
Kanagalakshmi Murugan \
International Journal on Science and Technology \
E-ISSN-2229-7677 \
License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijsat.org/ \
**Author Contact** \
**LinkedIn**: http://linkedin.com/in/kanagalakshmi-murugan-12b278199 | **Email**: kanagalakshmi2004@gmail.com

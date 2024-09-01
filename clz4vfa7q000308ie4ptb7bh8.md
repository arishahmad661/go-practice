---
title: "Step-by-Step Guide to Create a Snapshot, and Volume from Snapshot"
seoTitle: "Snapshot and Volume Creation Guide"
seoDescription: "Guide to creating AWS snapshots and volumes from snapshots with step-by-step instructions and screenshots"
datePublished: Sun Jul 28 2024 01:17:53 GMT+0000 (Coordinated Universal Time)
cuid: clz4vfa7q000308ie4ptb7bh8
slug: step-by-step-guide-to-create-a-snapshot-volume-from-snapshot-and-modify-a-volume
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722132661209/da2bb4fa-0df3-4cd6-947e-e64825b7ab33.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722132674466/8ec07a0a-0944-40ee-a44f-c2b6c190f44f.png
tags: aws, ebs, aws-ebs, ebs-snapshots, aws-ebs-beginners

---

### Creating a snapshot from volume

1. Open the **AWS Management Console** and navigate to the **EC2 dashboard**.
    
2. In the left navigation panel, choose **Volumes** under the **Elastic Block Store section.**
    
3. Select the volume you want to snapshot.
    
4. Got to **Action** -&gt; **Create snapshot**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722127923725/dd10f200-fea3-4750-baaa-df8e7cbfcd58.png align="center")
    
5. Write a suitable description that may be used to identify the snapshot.
    
6. ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722128022149/b975534d-4bf2-4e48-b071-1cffd21e4eab.png align="center")
    
    Scroll and click **Create Snapshot**.
    
7. In the left navigation panel, choose **Snapshots** under the **Elastic Block Store section.**
    
8. You can see the newly created snapshot.
    

### Creating volume from a snapshot

1. Select the snapshot.
    
2. Go to **Action** -&gt; **Create volume from snapshot**.
    
3. ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722128196054/4f78d5f3-f016-4f4a-9cba-bc881ec06fa6.png align="center")
    
    Select the **Volume type** to General Purpose SSD(GP2).
    
4. Select a **size** as 30 GiB.
    
5. Scroll and click **Create Volume**.
    
6. In the left navigation panel, choose **Volumes** under the **Elastic Block Store section.**
    
7. You can see the newly created volume.
---
title: "Step-by-Step Guide: Creating an Instance from a Snapshot"
seoTitle: "Create Instance from Snapshot Guide"
seoDescription: "Guide: Creating an Instance from an AWS Snapshot in 13 Steps"
datePublished: Sun Jul 28 2024 02:35:12 GMT+0000 (Coordinated Universal Time)
cuid: clz4y6pxc00040al3hk6mg1lz
slug: step-by-step-guide-creating-an-instance-from-a-snapshot
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722133928833/af4b26a5-6dbc-412c-8627-3ec5e16a8eda.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722134080211/6f4020c0-1f95-4bc8-9820-578c90bfbd46.png
tags: ec2, aws, aws-ec2, aws-ebs, aws-ami, aws-instance

---

1. Open the **AWS Management Console** and navigate to the **EC2 dashboard**.
    
2. In the left navigation panel, choose **Snapshots** under the **Elastic Block Store section.**
    
3. Select the snapshot from which you want to launch the instance.
    
4. Go to **Action** -&gt; **Create image from snapshot**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722133604819/f16222c7-6ed3-4de0-a307-f7ddc6516ece.png align="center")
    
5. Write a suitable image name.
    
6. Scroll down and click **Create image**.
    
7. In the left navigation panel, choose **AMIs** under the **Images.**
    
8. Select the newly created image.
    
9. Click on **Launch instance from AMI**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722133738767/5426a31d-4c72-4e8c-b08c-47e7a1f68e44.png align="center")
    
10. Write a suitable instance name.
    
11. Scroll down and select the **key pair**.
    
12. Click **Launch Instance**.
    
13. In the left navigation panel, choose **Instances** under the **Instances, here you can find the newly created instance.**
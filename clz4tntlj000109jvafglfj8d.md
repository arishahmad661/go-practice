---
title: "Step-by-Step Process to Detach and Modify EC2 Volumes in AWS"
seoTitle: "Detach and Modify EC2 Volumes Easily"
seoDescription: "Learn how to detach and modify EC2 volumes in AWS step-by-step"
datePublished: Sun Jul 28 2024 00:28:32 GMT+0000 (Coordinated Universal Time)
cuid: clz4tntlj000109jvafglfj8d
slug: step-by-step-process-to-detach-and-modify-ec2-volumes-in-aws
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722126132198/512cae08-a577-4471-a6bf-120251ad1b3a.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722126503198/2d1552e7-8ba7-408b-9c5f-c1ffa3b21063.png
tags: aws, ebs, aws-ebs, aws-ebs-introduction, aws-ebs-beginners

---

### Detach the volume

1. Open the **AWS Management Console** and navigate to the **EC2 dashboard**.
    
2. In the left navigation panel, choose **Volumes** under the **Elastic Block Store section.**
    
3. Select the volume you want to detach.
    
4. Go to **Actions** -&gt; **Detach volume**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722126393535/e295be2b-9a50-42a0-9d94-1c8f192b278d.png align="center")
    
5. Click **Detach**.
    

**Note**: To detach the root volume of an instance, you must first stop the instance. After detaching the root volume, the instance cannot be started again until the root volume is reattached.

### Modifying the volume

1. Select the volume you want to modify.
    
2. Go to **Actions** -&gt; **Modify volume**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722128944349/e679c449-a05c-4610-b35c-2c1b702929f2.png align="center")
    
3. Apply the changes you want, such as increasing the size from 30 to 35 GiB.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722128950231/3c147da5-e4bc-4c02-b973-6c056469dbcd.png align="center")
    
4. Scroll and click **Modify** -&gt; **Modify.**
    
5. You can check the size of the volume now it should be 35 GiB.
    

1. Select the volume you want to modify.
    
2. Go to **Actions** -&gt; **Modify volume**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722128944349/e679c449-a05c-4610-b35c-2c1b702929f2.png align="center")
    
3. Apply the changes you want, such as increasing the size from 30 to 35 GiB.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722128950231/3c147da5-e4bc-4c02-b973-6c056469dbcd.png align="center")
    
4. Scroll and click **Modify** -&gt; **Modify.**
    
5. You can check the size of the volume now it should be 35 GiB.
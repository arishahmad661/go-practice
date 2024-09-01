---
title: "Step by Step process of how to add and initialize EBS Volume on Windows EC2 Instance"
seoTitle: "Mount EBS Volume on Windows EC2 Instance"
seoDescription: "How to add and mount EBS Volume on Windows EC2 Instance step-by-step"
datePublished: Sun Jul 28 2024 00:09:59 GMT+0000 (Coordinated Universal Time)
cuid: clz4szyvn000509l2bj1x6s9r
slug: step-by-step-process-of-how-to-add-and-initialize-ebs-volume-on-windows-ec2-instance
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722125274539/23ea165a-47b0-4e89-abf5-170374489817.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722125371811/069f93a6-3f41-422f-9cfc-ea9c04df6d04.png
tags: aws, ebs, aws-ec2, aws-ebs, aws-ebs-beginners, amazon-ebs

---

### **Step 1: Creating a volume**

1. Open the AWS Management Console and navigate to the EC2 dashboard.
    
2. In the left navigation panel, choose **Volumes** under the **Elastic Block Store section.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722123713736/e5e5e28a-8e46-4b11-a4e7-9a3990cd8db3.png align="center")
    
3. Take note of existing Volumes.
    
4. Click on **Create Volume.**
    
5. Choose the **volume type** as General Purpose SSD (gp2).
    
6. Set **Size** to 30GiB.
    
7. Select the availability zome same as the availability zome as the instance with which you want to attach this volume.
    
8. Scroll down and click **Create Volume.**
    

### Step 2: Attach the volume to the instance

1. Select the volume, then click **Actions** -&gt; **Attach volume.**
    
2. ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722125063083/3aa67687-feda-4075-a5c9-be2bc1cea946.png align="center")
    
    Select the instance in which you want to add the volume.
    
3. Select the **device name** xvdb,if you are attaching the root volume then select /dev/sda1.
    
4. Click on **Attach volume**.
    

### Step 3: Connect to EC2 instance

1. In the left navigation panel, choose **Instances** under the **Instances** section.
    
2. Select the instance in which you have attached the volume.
    
3. Click on **Connect**.
    
4. Go to **RDP client.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722124088628/ef7f6e9d-ebaa-49b7-a373-af6e525b3057.png align="center")
    
5. Scroll and click on **Get password**.
    
6. Click on **Upload private key file** and Upload the private key you choose while creating this instance.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722124112928/0bf8ffe0-358f-4b43-9d96-74776b8dda86.png align="center")
    
7. Click on **Decrypt password**.
    
8. A password will appear under the RDP client copy this password.
    
9. Click on **Download remote desktop file.**
    
10. Open this **.rdp** file.
    
11. Click **Connect** -&gt; Paste the password -&gt; Click **OK** -&gt; Click **Yes.**
    
12. The remote desktop will open up.
    

### Step 4: Making the disk online and initializing it

1. Press **Ctrl + r**.
    
2. Type **diskmgmt.msc** and press enter.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722124311837/2d6d2b6b-89f7-43d6-b496-9f32eb18f081.png align="center")
    
3. Right-click on Disk 1 then click **Online**.
    
4. ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722124467139/f3783ec0-141c-4b48-9a3d-e4c2ae0769f2.png align="center")
    
    Right-click on Disk 1 again then click **Initialize Disk** -&gt; **OK**. If this disk is already initialized then this step will be skipped.
    
5. Now right-click on unallocated then click **New Simple volume** -&gt; **Next** -&gt; **Next** -&gt; **Next** -&gt; **Next** -&gt; **Finish**.
    
6. Open **File Explorer** then **This PC**.
    
7. You check a new drive **New Volume(D:)**.
    
8. You can create a text file on this drive to check its working.
    

**NOTE:** You can also add volume to the instance during the creation of instance by using Add Volume under **Configure storage.**

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722125217149/5852fbec-b4fc-4aa5-873d-f3e1cb590604.png align="center")
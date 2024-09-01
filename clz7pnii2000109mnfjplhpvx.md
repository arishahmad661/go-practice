---
title: "How to Create an Amazon S3 Bucket: Step-by-Step Guide"
seoTitle: "Create an Amazon S3 Bucket: Quick Guide"
seoDescription: "Learn how to create an Amazon S3 bucket step-by-step, ensuring secure data storage and easy accessibility"
datePublished: Tue Jul 30 2024 00:59:38 GMT+0000 (Coordinated Universal Time)
cuid: clz7pnii2000109mnfjplhpvx
slug: how-to-create-an-amazon-s3-bucket-step-by-step-guide
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722301055364/7d441384-4b0c-46c5-95ad-68e4fdfad9d6.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722301126247/f293660c-ad10-45cf-ba81-8f1cb6071a95.png
tags: aws, s3, aws-s3, s3-bucket, aws-s3-for-beginners

---

### Introduction

1. S3 stands for Simple Storage Service.
    
2. Amazon S3 is an object storage service that offers industry-leading scalability, data availability, security, and performance.
    
3. Store and protect any amount of data for a range of use cases, such as data lakes, websites, cloud-native applications, backups, archive, machine learning, and analytics.
    
4. Amazon S3 is designed for 99.999999999% (11 9's) of durability, and stores data for millions of customers all around the world.
    

### Steps to Create an S3 Bucket

1. Open the **AWS Management Console** and navigate to the **S3 dashboard**.
    
2. Click **Create bucket**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722299964480/fb19aa61-1dc5-4b42-86eb-23f7ce1980f4.png align="center")
    
3. Select **General purpose** bucket type.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722300008769/9a6f2b7a-ae48-4fe0-8d71-6bec71d585e1.png align="center")
    
4. Write a bucket name, this name must be globally unique.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297511455/5a02b511-e7c3-4a42-b05b-34d15614f746.png align="center")
    
5. Select **ACLs disabled** under **Object Ownership.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297537994/b7a5f094-c8e8-4941-bfd5-357618f3d746.png align="center")
    
6. Leave the **Block all public access** checked**.**
    
7. Set **Bucket Versioning** disabled**.**
    
8. Click **Create bucket**.
    
9. Open the newly created bucket by clicking on its name.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297808074/887d39dd-a1c3-4596-a1fd-ac268c6534c6.png align="center")
    
10. Click **Upload**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297830670/adb55fb1-0c38-438d-a8b1-3aa3374511f0.png align="center")
    
11. Now click on Add files.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297908287/836cfbcf-e114-4e23-86f5-c4bd71555c9e.png align="center")
    
12. Select any file you want to upload, preferably a text file.
    
13. Scroll and click **Upload**.
    
14. Select the file -&gt; **Actions** -&gt; **Download As** -&gt; Click on the file name, A new tab will open and download will start -&gt; **Close.**
    
15. Now Select the file -&gt; Click **Copy URL**. -&gt; Paste in a new browser tab, Unfortunately, you won't be able to view the file.
    
16. Now Select the file -&gt; Click **Open**, File will open up in a new tab.
    
17. You can see the file can be accessed only by this account and no one else can access it.
    
18. To make the Object's URL publically available -&gt; Select the bucket -&gt; Make public using ACL, this option must be disabled because we disabled ACLs in step 6.
    
19. To change it Go to permission.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722300158363/3436b808-1081-4d88-9d8b-fcf9ab98bf8c.png align="center")
    
20. Click **Edit** on **Block public access (bucket settings) -&gt;** Uncheck **Block all public access -&gt; Save changes -&gt;** type **confirm -&gt; Confirm.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722299478090/7c8695ff-7ba0-4715-824e-1db410f3645e.png align="center")
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722299495279/eca32f6b-d72c-4e77-9c02-3338a0630682.png align="center")
    
21. Click **Edit** on **Bucket Ownership** -&gt; **ACLs enabled** -&gt; Check the acknowledgement -&gt; Save changes.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722299346209/1b23e0ab-ddbc-489f-973b-4f46801f7c10.png align="center")
    
22. Now go back to **Objects**.
    
23. Select the object -&gt; **Actions** -&gt; **Make public ACL** -&gt; **Make Public**.
    
24. Now select the object -&gt; **Copy URL** -&gt; Paste it into another tab, and the file should open.
    
25. To delete the object, select the object -&gt;Click **Delete** -&gt; type permanently delete -&gt; Click **Delete Objects.**
    
26. To delete the bucket, select the bucket -&gt; Click **Delete** -&gt; type the bucket name -&gt; Click **Delete bucket**.
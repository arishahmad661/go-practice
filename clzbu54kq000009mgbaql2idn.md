---
title: "Setting Up Amazon S3 Replication Rules"
seoTitle: "Amazon S3 Replication Guide"
seoDescription: "Step-by-step guide to set up and manage Amazon S3 replication rules between buckets"
datePublished: Thu Aug 01 2024 22:16:23 GMT+0000 (Coordinated Universal Time)
cuid: clzbu54kq000009mgbaql2idn
slug: setting-up-amazon-s3-replication-rules
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722550400270/4f133ead-ac36-4a1f-896c-eddff3b770b7.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722550561618/de7d4672-9559-4f89-b8bd-d74900a2b2d6.png
tags: aws, s3, aws-s3, aws-s3-versioning, s3-replication, s3-bucket, aws-s3-for-beginners, s3replication

---

1. Open the **AWS Management Console** and navigate to the **S3 dashboard**.
    
2. Click **Create bucket**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722299964480/fb19aa61-1dc5-4b42-86eb-23f7ce1980f4.png?auto=compress,format&format=webp&auto=compress,format&format=webp align="left")
    
3. Select **General purpose** bucket type.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722300008769/9a6f2b7a-ae48-4fe0-8d71-6bec71d585e1.png?auto=compress,format&format=webp&auto=compress,format&format=webp align="left")
    
4. Write a bucket name, this name must be globally unique.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297511455/5a02b511-e7c3-4a42-b05b-34d15614f746.png?auto=compress,format&format=webp&auto=compress,format&format=webp align="left")
    
5. Select **ACLs enabled Object Ownership.**
    
6. Uncheck **Block all public access -&gt; Check the Acknowledgement.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722302267197/d2338c6a-6027-482e-b590-291584bfe4c2.png?auto=compress,format&format=webp align="left")
    
7. Enable **Bucket Versioning**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722302343722/a08e7765-cdae-4a06-8010-486562292af1.png?auto=compress,format&format=webp align="left")
    
8. Click **Create bucket**.
    
9. Create one more bucket with the same configuration.
    
    One bucket will act as the source and the other as the destination bucket.
    
10. Open any bucket by clicking on its name. This bucket will act as a source bucket.
    
11. Click on **Management**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722548155495/0a880881-a9f6-4426-b0ae-d5050db803c2.png align="center")
    
12. Click **Create replication rule.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722549027641/c59839ab-7a1f-4ea7-90e2-aeecdb920476.png align="center")
    
13. Enter a suitable replication name.
    
14. Scroll and choose **Apply to all objects in the bucket** under **Choose a rule scope**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722549075307/47367bc7-9e9a-4384-a0a9-12d763a17018.png align="center")
    
15. Click **Browse S3**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722549098144/e331792c-78d9-4383-b1b8-a74e724676dc.png align="center")
    
16. Choose the destination bucket.
    
17. Click **Choose path.**
    
18. Select **Create new role** under **IAM role**.
    
19. Click **Save**.
    
20. Select any under **Replicate existing objects?** objects click **Submit**.
    
21. Open the source bucket and upload a file.
    
22. Open the destination bucket, and you'll find the uploaded file there as well.
    
23. Go to the source bucket.
    
24. Select the bucket -&gt; **Copy URL** -&gt; Paste the URL in the other browser tab. You'll not be able to access the file. Try the same in the destination bucket too, you'll get the same result.
    
25. In the source bucket, Select the object, click on **Action** -&gt; **Make public using ACL** -&gt; **Make public**.
    
26. Now access the bucket. Select the bucket -&gt; **Copy URL** -&gt; Paste the URL in other browser tab, you'll see the contents of the file.
    
27. Navigate to the destination bucket, you'll be able to access the file there too.
    
28. Now, permanently delete the file in the source bucket.
    
    **Note**: Make sure **Show versions** is enabled when you delete; otherwise, the file will not be permanently deleted.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722550065773/74f57f94-8d1c-42ac-b9d2-900c12379649.png align="center")
    
29. Now to go destination bucket, you'll see the file in the destination bucket is not deleted.
    
30. To delete the replication rule, open the source bucket, click **Management**, select the replication rule you want to delete under **Replication rules**, then click **Delete** and confirm by clicking **Delete replication rule**.
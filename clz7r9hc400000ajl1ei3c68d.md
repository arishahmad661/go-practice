---
title: "Understanding Versioning for Amazon S3 Buckets"
seoTitle: "Amazon S3 Bucket Versioning Explained"
seoDescription: "Learn how to enable and manage versioning for Amazon S3 buckets with step-by-step instructions"
datePublished: Tue Jul 30 2024 01:44:43 GMT+0000 (Coordinated Universal Time)
cuid: clz7r9hc400000ajl1ei3c68d
slug: understanding-versioning-for-amazon-s3-buckets
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722303738910/fa4de092-7c74-4578-8d46-0276ccb92c48.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722303867832/fffce32e-15ac-462a-9b99-8c6ac7effd0e.png
tags: aws, versioning, s3, aws-s3, aws-s3-versioning, s3-bucket, aws-s3-for-beginners

---

1. Open the **AWS Management Console** and navigate to the **S3 dashboard**.
    
2. Click **Create bucket**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722299964480/fb19aa61-1dc5-4b42-86eb-23f7ce1980f4.png?auto=compress,format&format=webp align="left")
    
3. Select **General purpose** bucket type.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722300008769/9a6f2b7a-ae48-4fe0-8d71-6bec71d585e1.png?auto=compress,format&format=webp align="left")
    
4. Write a bucket name, this name must be globally unique.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297511455/5a02b511-e7c3-4a42-b05b-34d15614f746.png?auto=compress,format&format=webp align="left")
    
5. Select **ACLs enabledObject Ownership.**
    
6. Uncheck **Block all public access -&gt; Check the Acknowledgement.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722302267197/d2338c6a-6027-482e-b590-291584bfe4c2.png align="center")
    
7. Enable **Bucket Versioning**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722302343722/a08e7765-cdae-4a06-8010-486562292af1.png align="center")
    
8. Click **Create bucket**.
    
9. Open the newly created bucket by clicking on its name.
    
10. Click **Upload**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297830670/adb55fb1-0c38-438d-a8b1-3aa3374511f0.png?auto=compress,format&format=webp align="left")
    
11. Now click on **Add files.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722297908287/836cfbcf-e114-4e23-86f5-c4bd71555c9e.png?auto=compress,format&format=webp align="left")
    
12. Select any file you want to upload, preferably a text file.
    
13. Scroll and click **Upload**.
    
14. Select the object -&gt; **Actions** -&gt; **Make public ACL** -&gt; **Make Public**.
    
15. Now select the object -&gt; **Copy URL** -&gt; Paste it into another tab, and the file should open.
    
16. Select the file -&gt; **Delete** -&gt; type "delete" -&gt; Click **Delete objects**.
    
17. Now enable **Show versions**. Now you'll be able to see the deleted text file and the Delete Marker. If this delete marker is deleted then the text file can be recovered.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722302633360/36c66a95-089a-4d3e-8870-ec756b27980e.png align="center")
    
18. Select the delete marker -&gt; **Delete** -&gt; type "permanently delete" -&gt; Click **Delete Objects**.
    
19. Disable the Show version. You will see the text file is recovered. To delete the file permanently, delete the text file while Show versions is enabled before deleting the marker.
    
    **Note**: Deleting anything while Show versions is enabled will result in permanent deletion.
    
20. Now open this text file on your computer and make a few changes to it.
    
21. Now upload the file again. The file will be updated.
    
22. Select the object -&gt; **Actions** -&gt; **Make public ACL** -&gt; **Make Public**.
    
23. Now select the object -&gt; **Copy URL** -&gt; Paste it into another tab, and the file should open. You can see the changes made.
    
24. Now enable the Show version, and you'll see the same file name two times, the one with 'L' is the old file and the other one is the new file.
    
25. Delete the file without 'L'.
    
26. Disable the show version.
    
27. Now select the object -&gt; **Copy URL** -&gt; Paste it into another tab, and the file should open. You'll see that the older version of the file.
    
28. To delete the object, select the object -&gt;Click **Delete** -&gt; type permanently delete -&gt; Click **Delete Objects.**
    
29. To delete the bucket, select the bucket -&gt; Click **Delete** -&gt; type the bucket name -&gt; Click **Delete bucket**.
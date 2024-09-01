---
title: "How to Easily Create and Manage IAM User Groups in AWS"
seoTitle: "Creating and Managing AWS IAM User Groups"
seoDescription: "Learn to create and manage IAM user groups in AWS effortlessly. Guide includes adding users, assigning policies, and copying permissions"
datePublished: Sun Aug 04 2024 23:18:56 GMT+0000 (Coordinated Universal Time)
cuid: clzg6p4dm000009l51cbjbg31
slug: how-to-easily-create-and-manage-iam-user-groups-in-aws
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722813366608/0945a863-774f-4bd8-b8c4-77e9952c1d7e.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722813518172/74f446ef-3e4f-4656-b060-2839693056f7.png
tags: aws, iam, aws-iam, aws-iam-policies, iam-users, iam-user-groups

---

## Adding users to a User Group

1. Open the **AWS Management Console** and navigate to the **Identity and Access Management (IAM).**
    
2. Create 2-3 users, you can use [Creating And Assigning Policy to IAM Users](https://arishahmad.hashnode.dev/aws-creating-and-assigning-policy-to-iam-users).
    
3. Click on **User groups** from the left navigation panel.
    
4. Click **Create group**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812333763/6ee6dd63-caf3-44f0-86e8-64fd9da1e17c.png align="center")
    
5. Write a suitable name.
    
6. You can select all the users you want to add to this group.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812358495/71733ed6-c9bc-4559-892a-56ebb6acb86a.png align="center")
    
7. Search and check the policies you want to assign (Eg. **AmazonEC2ReadOnlyAccess**).
    
8. Click **Create user group**.
    
9. Now open the user and go to **Permissions**, you can check the policies given under **Permissions policies.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812380432/2597da21-fce9-4769-bfdd-faeebe015057.png align="center")
    

## Adding users to multiple User Groups

1. Create another user group and add the same users that were added in the previous user group.
    
2. Search and select some other policies from the previous one (Eg. **AmazonEC2ReadOnlyAccess**).
    
3. Now open the user and go to **Permissions**, you can check the policies from different user groups given under **Permissions policies.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812409438/e211b6ec-fff1-4d4d-af74-e1ef9763f467.png align="center")
    

## Giving permissions to a specific user

1. Open the user.
    
2. Under **Permissions,** click **Add permissions** -&gt; **Add permissions**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812433689/89bd89b7-593f-457b-9d76-f9052af9933a.png align="center")
    
3. Select **Attach policies directly**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812449465/74ed4571-de52-40dd-8fa5-76e12c2906a9.png align="center")
    
4. Search and select any policy (Eg. **AmazonDynamoDBFullAccess**).
    
5. Click **Next** -&gt; **Add permissions**.
    

To check, open the user under **Permissions** and check the newly added permission.

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722812475239/01cfc8cb-7782-4221-a9f1-5bb76b6a1dab.png align="center")

## Copy permission to a user

1. Create a new user.
    
2. Open this user.
    
3. Under **Permissions,** click **Add permissions** -&gt; **Add permissions**.
    
4. Click Copy permissions.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722813238684/594b0bb3-e7b4-4ca4-a070-a3c7bd3e6fb3.png align="center")
    
5. Select any user whose permissions are needed to be copied.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722813256207/ee6a345a-f9fb-4ed8-a801-73c1292a5b8a.png align="center")
    
6. Click **Next** -&gt; **Add permissions**.
    
7. Now open this user, and click **Permissions**, under **Permissions policies** you can check all the permissions given to this user which were given to the user from which the permissions are copied.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722813302640/67fb3856-c01d-4f43-9219-14d8f2fd3698.png align="center")
---
title: "How to Create and Attach IAM Roles in AWS"
seoTitle: "Attaching IAM Roles in AWS"
seoDescription: "Learn to create and attach IAM roles in AWS, including detailed steps and permissions setup for EC2 services"
datePublished: Mon Aug 05 2024 02:40:42 GMT+0000 (Coordinated Universal Time)
cuid: clzgdwllh000p09mi7zo751s9
slug: how-to-create-and-attach-iam-roles-in-aws
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722825398065/07dd9d8c-35e0-42d5-ba74-fb2b8352683b.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722825609759/4dc9eff4-96e6-4919-8f12-70dc899efa34.png
tags: aws, iam, aws-iam, iam-role, iam-role-in-aws, aws-iam-role

---

### Introduction

An IAM *role* is an IAM identity that you can create in your account that has specific permissions. An IAM role is similar to an IAM user, in that it is an AWS identity with permission policies that determine what the identity can and cannot do in AWS. However, instead of being uniquely associated with one person, a role is intended to be assumable by anyone who needs it. Also, a role does not have standard long-term credentials such as a password or access keys associated with it. Instead, when you assume a role, it provides you with temporary security credentials for your role session.

For more details [visit](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html).

### Creating IAM Roles

1. Open the **AWS Management Console** and navigate to the **Identity and Access Management (IAM).**
    
2. Click on **Roles** from the left navigation panel.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722824606774/4e2009dd-12a6-4030-9935-6cdf05d908b8.png align="center")
    
3. Click **Create role**.
    
4. Select **AWS Service**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722824638146/6a4ff80a-8389-4cf4-91af-cd44e3b5660b.png align="center")
    
5. Select **EC2** in the **Use case.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722824661423/d4901825-a785-4628-8528-bc46e40b7cb0.png align="center")
    
6. Click **Next**.
    
7. Search **AmazonS3FullAccess** in **Permissions policies, and select it.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722824691317/77e82a8f-a6e2-4502-856e-6c2c12cba0ba.png align="center")
    
8. Write a suitable role name.
    
9. ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722824706169/e8330327-87fa-4964-b323-0337a75e092b.png align="center")
    
    Click **Next** -&gt; **Create Role**.
    
10. Open the role by clicking on its name.
    
11. Select **Permissions**, under **Permissions policies** you can check the policy names.
    

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722824759891/fd068280-1b30-4a28-b11b-6812e8fd441f.png align="center")

### Attaching role to a service

1. Create an EC2 instance.
    
2. Select the instance.
    
3. Go to **Actions** -&gt; **Security** -&gt; **Modify IAM role**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722825352973/75091ce0-596a-4e02-bdd9-bb08c3c5d0d4.png align="center")
    
4. Select the IAM role from the from down.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722825378456/663ac508-39b6-4274-a302-f485a2f81c26.png align="center")
    
5. Click **Update IAM role**.
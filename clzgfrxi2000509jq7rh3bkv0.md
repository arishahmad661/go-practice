---
title: "How to Create and Attach IAM Policies in AWS"
seoTitle: "Attach AWS IAM Policies: Step-by-Step Guide"
seoDescription: "Step-by-step guide on creating and attaching IAM policies in AWS for managing user and resource permissions"
datePublished: Mon Aug 05 2024 03:33:03 GMT+0000 (Coordinated Universal Time)
cuid: clzgfrxi2000509jq7rh3bkv0
slug: how-to-create-and-attach-iam-policies-in-aws
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722828506644/3acb085a-7ded-4b0b-9b87-bec5d68b81b7.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722828762114/815cf150-c232-4a69-ae43-13d62b62164e.png
tags: aws, iam, aws-iam, aws-iam-policies, iam-policies, iam-policy, aws-iam-policy

---

### Introduction

You manage access in AWS by creating policies and attaching them to IAM identities (users, groups of users, or roles) or AWS resources. A policy is an object in AWS that, when associated with an identity or resource, defines their permissions. AWS evaluates these policies when an IAM principal (user or role) makes a request. Permissions in the policies determine whether the request is allowed or denied. Most policies are stored in AWS as JSON documents. AWS supports six types of policies: identity-based policies, resource-based policies, permissions boundaries, Organizations SCPs, ACLs, and session policies.

For more details [visit](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).

### Create IAM Policies

1. Open the **AWS Management Console** and navigate to the **Identity and Access Management (IAM).**
    
2. Click on **Policies** from the left navigation panel.
    
3. Click **Create policy**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826458792/d08e9f9a-cf6e-440b-8c7d-2aa8a2c55b40.png align="center")
    
4. Choose **S3** from the down under **Service**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826474942/967d6586-1c6a-4c9a-9b97-efcfadda9baa.png align="center")
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826488233/b8764485-b627-4a5b-b4d8-9a0d7d8438eb.png align="center")
    
5. Click **All list actions** under **Access level.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826504259/11aff54a-4887-44f0-aa4b-49eead7392a0.png align="center")
    
6. Select **All** under **Resources**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826520183/568aeef7-b07e-44d9-809a-ec37715aef65.png align="center")
    
7. Click **Next**.
    
8. Write a suitable policy name.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826540857/bdf3d128-f3e1-4979-94c6-337dc9ae576d.png align="center")
    
9. Click Create policy.
    

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722826558541/3db5d935-8ab0-4972-b2fc-4caced5a336f.png align="center")

### Creating and attaching policy with an IAM user

1. Click on **Users** from the left navigation panel.
    
2. Click **Create user**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722828318677/a6131b2b-1823-4173-aaa1-2b0da41bee47.png align="center")
    
3. Write a suitable user name.
    
4. Check **Provide user access to the AWS Management Console**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722828359677/749c2846-2cce-4f44-bcdf-c9138804c98f.png align="center")
    
5. Select **I want to create an IAM user**.
    
6. Click **Custom password** and type a suitable password.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722828415256/25647c05-e880-455d-847a-3d6d52168255.png align="center")
    
7. Uncheck **Users must create a new password at next sign-in** for now.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722828437402/8458ead9-add2-4c42-8780-daf2d717b322.png align="center")
    
8. Click **Next**.
    
9. Select **Attach policies directly**.
    
10. Search and select the newly created policy.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722828482318/6dae53ac-9568-4afa-916c-117bafcdad52.png align="center")
    
11. Click **Next** -&gt; **Create user**.
    
12. Create an S3 bucket.
    
13. Login as an IAM user with newly created credentials.
    
14. The S3 bucket can be viewed by this IAM user.
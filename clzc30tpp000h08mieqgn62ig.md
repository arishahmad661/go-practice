---
title: "AWS: Creating And Assigning Policy to IAM Users"
seoTitle: "AWS IAM User Management Guide"
seoDescription: "Guide to creating and managing IAM users in AWS. Step-by-step instructions for setting up user access and permissions"
datePublished: Fri Aug 02 2024 02:24:59 GMT+0000 (Coordinated Universal Time)
cuid: clzc30tpp000h08mieqgn62ig
slug: aws-creating-and-assigning-policy-to-iam-users
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1722563908893/2c7fd8b7-4408-4551-8165-cd96d96b9a4f.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1722563964095/a6fbd2bc-47c2-49a1-93bd-4671957f9867.png
tags: aws, iam, aws-iam, aws-iam-policies, iam-identities, aws-iam-identity-center, aws-iam-permissions

---

1. Open the **AWS Management Console** and navigate to the **Identity and Access Management (IAM).**
    
2. Click on **Users** from the left navigation panel.
    
3. Click **Create user**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722563141861/a440cfd9-f8ee-4cab-beb0-13b0c0872b71.png align="center")
    
4. Assign a suitable name to the user.
    
5. Select **Provide user access to the AWS Management Console**.
    
6. Select **I want to create an IAM user.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722563169649/4008815b-2ce6-42ff-9213-41938128ca6d.png align="center")
    
7. Select **Custom password** and enter a password fulfilling the criteria.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722563206305/76f83cbe-7907-418a-98c5-20d18b1300d0.png align="center")
    
8. Uncheck **Users must create a new password at next sign-in** for now.
    
9. Click **Next**.
    
10. Select **Attach policies directly**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722563238564/81c17290-e71b-46e0-8231-a449bf8e2a51.png align="center")
    
11. Search **AmazonEC2ReadOnlyAccess** and select it.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722563281367/5147ecb0-ad8f-487c-b742-3c118699595f.png align="center")
    
12. Click **Next**.
    
13. Scroll and click **Create user**.
    
14. Now create an EC2 instance and log in as an IAM user using the user name and password just created to check the newly created EC2 instances.
    
15. To delete a user, select the user -&gt; **Delete** -&gt; type the user name -&gt; **Delete user**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722564067529/126ad059-d5c7-4d7b-8652-c7f926aecc0f.png align="center")
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722564083462/c240d68c-8746-4829-a4eb-65209767c69c.png align="center")
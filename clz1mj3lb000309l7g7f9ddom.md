---
title: "Hosting Your Website: Launching AWS EC2 Instance with Windows"
seoTitle: "Amazon EC2 Windows Website Hosting Guide"
seoDescription: "Learn how to host your website on Amazon EC2 with Windows in this step-by-step guide"
datePublished: Thu Jul 25 2024 18:45:36 GMT+0000 (Coordinated Universal Time)
cuid: clz1mj3lb000309l7g7f9ddom
slug: hosting-your-website-on-amazon-ec2-with-windows-a-complete-guide
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1721932911980/e0b249a6-afdf-4c50-904e-efe564b00ea2.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1721933122432/c707655d-a67c-463b-bd2b-0245f9f719ca.png
tags: ec2, aws, aws-ec2, ec2-instance-types, ec2-instance, aws-instance-connect, aws-connect, aws-instance, ec2-hosting, ec2-windows

---

1. Open the AWS Console using the link [**https://aws.amazon.com/console/**](https://aws.amazon.com/console/)**.**
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*DikCh2PHRsw_FSQYODq9gw.png align="left")
    
2. Click on **Sign In to the Console**.
    
3. After signing up search **EC2** in the search bar.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*QKdZNf_d0vzYBCQoYkVLhA.png align="left")
    
4. Click on **EC2**. This window will open up.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*5QS_IBQE-xNKi0jO1eup-Q.png align="left")
    
5. Scroll down and Click on **Launch Instances**.
    
6. Give a suitable name for the instance.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*UG9_gQucfudCz5eLkdpngQ.png align="left")
    
7. Select **Windows under Application and OS Images (Amazon Machine Image)**.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*2KBS7-utMovE8viVUcbf5w.png align="left")
    
8. Scroll down to the **Key pair**.
    
9. Click on **Create new key pair**.
    
10. Enter the key name, Select Key pair type **RSA** click on **.pem** under Private key file format then click on **Create key pair**.
    
    ![](https://miro.medium.com/v2/resize:fit:588/1*_DOg4jN47l3SzjrZX9PvHw.png align="left")
    
11. The key file will start downloading.
    
12. Enable **Allow HTTPS traffic from the internet**, **Allow HTTP traffic from the internet**, and **Allow RDP traffic from** under network setting.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*1USoNdCcvJK1YbDpYP1GAg.png align="left")
    
13. Scroll down and click on **Launch Instance**.
    
14. Click on the **instance** from the left panel of the EC2 window.
    
15. Wait for 4–5 mins until the instance state is **notrunning** and the Status check is **not 2/2 checks passed**.
    
16. Select the instance then click on **Connect**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722021382739/28337b34-e7b3-42cd-bd37-4b5d201e9e55.png align="center")
    
17. Go to **RDP client**.
    
18. Click on **Download remote desktop file**.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*aH4v4gpEJ_tJjyhMnO3oMw.png align="left")
    
19. Then scroll down and click on **Get Password**.
    
20. Click on **Upload private key file**.
    
21. Then upload the key pair file that was created previously.
    
    ![](https://miro.medium.com/v2/resize:fit:700/1*6kDxiJ4WBXVWVxeR6pUKkg.png align="left")
    
22. Scroll down and click on **Decrypt password**.
    
23. A password will appear under the RDP client.
    
24. Copy this password.
    
25. Open the **.rdp** file downloaded previously.
    
26. Click **connect**, paste the **password**, and click **yes**. The remote desktop will open up.
    
27. Click on the **start button**.
    
28. Open **Server Manager**.
    
29. Wait for 4–5 minutes until the setup is complete.
    
30. Click on **add roles and features**.
    
31. Click on the **next** 3 times.
    
32. Now, scroll down and select **Web Server (IIS)**, then click on Add Feature.
    
33. Now click the **next** 4 times and then **install**.
    
34. Wait until the installation is finished.
    
35. Close the server manager when installation is complete.
    
36. Open **File Explorer** and go to **This PC**.
    
37. Open **Local Disk(C:)** -&gt; **inetpub** -&gt; **wwwroot**.
    
38. Delete all files present in **wwwroot**.
    
39. Create a **text file** and write something on it.
    
40. Press **Ctrl + Shift + S**, then enter the file name as **index.html** and select **All Files** from the **Save as type** dropdown. Then click save.
    
41. Go back to the instance on the AWS console.
    
42. Copy the **Public IPv4 address** of the instance.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722021441028/3acde0fd-85c0-4e9d-89f8-be9269b6e9d1.png align="center")
    
43. Paste it on the browser, and you will see the text you wrote in the **index.html** file.
    
44. To terminate the instance select the instance click on **instance state** and then click on **terminate**.
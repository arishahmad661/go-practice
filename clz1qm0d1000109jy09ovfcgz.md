---
title: "Create Amazon Machine Images (AMIs) from EC2 Instances: An Easy Guide"
seoTitle: "Create AMIs from EC2 Instances: Easy Guide"
seoDescription: "Step-by-step guide to create Amazon Machine Images (AMIs) from AWS EC2 instances"
datePublished: Thu Jul 25 2024 20:39:50 GMT+0000 (Coordinated Universal Time)
cuid: clz1qm0d1000109jy09ovfcgz
slug: create-amazon-machine-images-amis-from-ec2-instances-an-easy-guide
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1721939577112/b36be0a0-c51d-44b8-9ac6-416ce9c65df8.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1721939737741/0fce804e-916d-406e-bfe6-b7b3ab01d04d.png
tags: ec2, aws, ec2-instance-types, aws-ami, ec2-instance, aws-instance-connect, aws-instance, ec2-hosting

---

An Amazon Machine Image (AMI) is an image provided by AWS that provides the information required to launch an instance. You must specify an AMI when you launch an instance. You can launch multiple instances from a single AMI when you require multiple instances with the same configuration. You can use different AMIs to launch instances when you require instances with different configurations.

## Steps to Create an AMI

1. Create an EC2 instance. You can check out an [AWS EC2 Instance with Windows](https://arishahmad.hashnode.dev/hosting-your-website-on-amazon-ec2-with-windows-a-complete-guide) or an [AWS EC2 Instance with Linux](https://arishahmad.hashnode.dev/hosting-your-website-on-amazon-ec2-with-windows-a-complete-guide).
    
2. Now select the instance and click on **Actions**.
    
3. Select **Image and Templates**.
    
4. Choose **Create Image**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1721938119304/d6015945-f0ed-48a4-9334-60d319d790c1.png align="center")
    
5. Write down a suitable name for this image.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1721938211667/8d76e287-b416-4bae-a98b-495b88f98129.png align="center")
    
6. Scroll down and click **Create Image**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1721938248533/5c069af7-3f82-435e-997a-3781edb09ab5.png align="center")
    
7. Select AMIs under Images from the left panel.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1721938297937/818993d3-57f5-4045-9195-4c20bb6c5e26.png align="center")
    
8. Wait until the image of status is not **Available**.
    
    **NOTE:** Do not terminate the instance before the image is created; otherwise, the image will not be generated. Once the image status is available, you can terminate the instance.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1721938359789/62939b5a-ac25-408d-b957-c3ace9e94541.png align="center")
    
9. Select this image, and go to **Actions**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1722021838963/2cc393ef-f57f-4711-82f0-c3d28b8d444c.png align="center")
    
10. Choose **Deregister AMI** and again **Deregister AMI** to delete the AMI.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1721938496191/cbe7dfbd-7f36-4171-ae2f-d116ea369d7f.png align="center")
    
11. Now **Terminate** the instance.
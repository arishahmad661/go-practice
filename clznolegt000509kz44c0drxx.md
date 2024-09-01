---
title: "Accessing S3 Bucket Images with API"
seoTitle: "API Access for S3 Bucket Images"
seoDescription: "Guide to setting up an API for accessing Amazon S3 bucket images using AWS Lambda and API Gateway, with testing via Postman"
datePublished: Sat Aug 10 2024 05:14:19 GMT+0000 (Coordinated Universal Time)
cuid: clznolegt000509kz44c0drxx
slug: accessing-s3-bucket-images-with-api
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1723266679417/9b63f8b9-3601-4b6c-a570-6051737cbc8e.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1723266829362/3dd1f5c0-94ec-4bea-905f-4af94b92a2e2.png
tags: postman, lambda, aws, rest-api, aws-lambda, s3, api-gateway, lambda-function, aws-s3, aws-apigateway, aws-projects, s3-bucket, aws-lambda-for-beginners, aws-s3-for-beginners, aws-api-gateway

---

## Introduction

This blog will guide you through setting up an API to access images stored in Amazon S3. We’ll leverage AWS Lambda to create a serverless function that handles image retrieval, and API Gateway to expose this functionality as a RESTful API. To ensure everything works as expected, we'll use Postman for testing and validation.

## Step 1: Setting Up The S3 Bucket

1. Open the **AWS Management Console** and navigate to the **S3.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723180942878/27d7d370-e3a8-4062-b07f-88f5951a019b.png?auto=compress,format&format=webp align="left")
    
2. Click **Create Bucket**.
    
3. Give it a suitable name.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723263985617/dfadb10e-1cc8-4396-8af5-27ab49fa8e17.png align="center")
    
4. Scroll and click **Create Bucket**.
    
5. Open the bucket.
    
6. Click Upload.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264041311/c9d071fc-fd55-4c7e-ae45-0f8cf76663a4.png align="center")
    
7. Click **Add files**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264083021/3c193559-6457-47f5-8074-242e487661de.png align="center")
    
8. Upload an images.
    
    Make sure to use smaller images (in kb) for faster access.
    

## Step 2: Creating The Lambda Function

1. Open the **Lambda** and in **AWS Management Console.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264097459/1268338f-75c2-4d6a-bb27-c553045be2af.png align="center")
    
2. Click **Create a function**.
    
3. Give a suitable function name.
    
4. Select **Python 3.8** in **Runtime**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264127443/54089111-bce1-4d52-b20c-2dc744d99ee0.png align="center")
    
5. In **Advanced settings** enable **Enable function URL**, select **NONE** in **Auth type,** and enable **Configure cross-origin resource sharing (CORS)**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264144520/f61ff290-f8de-4cf4-b16a-494d87a0b565.png align="center")
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264181603/4fb70275-6a51-426f-8df9-9728c9ce3daf.png align="center")
    
6. Click on the **Create function**.
    
7. Open the function.
    
8. In the code area paste this code, save it, and click Deploy.
    
    ```bash
    import base64
    import boto3
    
    s3 = boto3.client('s3')
    
    def lambda_handler(event, context):
        try:
            # Extract bucket name and file name from the event
            bucket_name = event["pathParameters"]["bucket"]
            file_name = event["queryStringParameters"]["file"]
            
            # Fetch the file from S3
            response = s3.get_object(Bucket=bucket_name, Key=file_name)
            file_content = response["Body"].read()
            
            # Return the image as a base64 encoded string
            return {
                "statusCode": 200,
                "headers": {
                    "Content-Type": "image/jpeg",  # Adjust this if you have different image formats
                    "Content-Disposition": f"inline; filename={file_name}"  # 'inline' displays the image in the browser
                },
                "body": base64.b64encode(file_content).decode('utf-8'),
                "isBase64Encoded": True
            }
            
        except s3.exceptions.NoSuchKey:
            # Handle the case where the file does not exist in the bucket
            return {
                "statusCode": 404,
                "body": f"File {file_name} not found in bucket {bucket_name}."
            }
        
        except Exception as e:
            # Handle any other exceptions
            return {
                "statusCode": 500,
                "body": f"An error occurred: {str(e)}"
            }
    ```
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264211654/16028040-d20d-4dea-9373-57e6bf37a40b.png align="center")
    
9. Now click on **Configuration** -&gt; **Permissions**.
    
10. Click on the Role name, and a new tab will open up.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264229561/481d0285-00be-4f87-a5e0-064ef79a79af.png align="center")
    
11. Scroll and click **Add permissions** -&gt; **Attach policies**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264263808/e9c0a788-9380-4a43-a80d-2a814054019f.png align="center")
    
12. Search and select **AmazonS3FullAccess** then click **Add permissions**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264279736/0c56da78-51c1-4ff6-8867-fb84dbd13560.png align="center")
    

## **Step 3: Setting Up API Gateway**

1. Open the **API Gateway** and in **AWS Management Console.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264300324/6a344f59-7c41-4599-8070-cda5dceb93a6.png align="center")
    
2. Scroll and click **Build** on **Rest API**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264318106/21e7104e-3a3d-4785-8390-a93e5d9d618b.png align="center")
    
3. Select **New API** and give a suitable name.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264339847/0a963424-2b3f-4bf8-8dbf-50e0982370dc.png align="center")
    
4. Click on **Create API**.
    
5. Open the API.
    
6. Click on **Create resource**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265524234/7ba502ec-b187-46fa-864f-1b0620edc499.png align="center")
    
7. Give a resource name **{bucket}**.
    
    **Note**: Do not remove the curly bracket it will hold the bucket name which will be sent with the API.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265576951/32ee20cf-5621-44b2-a286-50ca39cacb3f.png align="center")
    
8. Click **Create resource**.
    
9. Now click on the **Create method**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265650851/d5597e98-31c4-47e3-94a8-fcb499bd5481.png align="center")
    
10. Select **GET** for the method type and the **Lambda function** for the **integration type**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264525862/cc2831d7-b319-45d5-865a-0b80520ba114.png align="center")
    
11. You can enable **Lambda proxy integration and select the Lambda function from the drop-down.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264678355/1181ba64-43c0-4139-a802-7fe6552a39c2.png align="center")
    
12. In the **Method request settings**, select **Validate query string parameters and headers** under **Request validator.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264551987/d042fa3b-176b-409e-b678-784fb9c98eb1.png align="center")
    
13. In **URL query string parameters** click **Add query string**, type **file** under **Name,** and check **Required.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264742356/a8534593-cf46-4da2-af33-6975d3705674.png align="center")
    
14. Click **Create method.**
    
15. From the left navigation panel click **API settings.**
    
16. Click on **Manage media types**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723264779752/d9c4f24f-b5ad-4598-ac31-93d963174bed.png align="center")
    
17. Then click on **Add binary media types**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265738976/acbb939a-0f06-48fa-bbcb-db030a4282e0.png align="center")
    
18. Enter **\*/\*** and click on **Save changes**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265717362/fdf89db6-6555-456c-86fb-ea8354cd6b5e.png align="center")
    
19. From the left navigation panel click **Resources.**
    
20. Click on **Deploy API**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265945980/a071a1ef-8d32-41f9-bf0c-68aabe9296fd.png align="center")
    
21. Select **\*New Satge\*** and give a suitable name.
    
22. Click on **Deploy**.
    
23. From the left navigation panel click **Stages.**
    
24. Click on **+** until to see **GET** then click on **GET**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265977837/6e8dfcca-20be-4432-82c8-13ae80c0c7c0.png align="center")
    
25. Copy the **Invoke URL.**
    

## Step 4: Testing the API with Postman

1. Open **Postman**.
    
2. Paste the link and replace **{bucket}** with the you s3 bucket name and add **?file=file\_name**. Replace file\_name with the file name that you upload in the s3 bucket (**Example**: if the **URL** is **https://url.com**, the **s3 bucket name** is **images.xyz,** and the **file name** is **maths.jpg** then use the link **https://url.com/images.xyz?file=maths.jpg**)
    
3. Click Send and you can see the file in Body section.
    
    You can also use this link in a browser tab and fetch the image.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723266085056/c261f2e1-a8dc-43f4-aa0a-78c4e693b552.png align="center")
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723265193310/58f5b66b-c91f-4126-bd54-6b17cb343d93.png align="center")
---
title: "Hosting Your Serverless Website Using AWS"
seoTitle: "AWS Serverless Website Hosting Guide"
seoDescription: "Learn how to host a serverless website on AWS with S3, CloudFront, Route53, DynamoDB, and Lambda in this step-by-step guide"
datePublished: Fri Aug 09 2024 06:56:09 GMT+0000 (Coordinated Universal Time)
cuid: clzmcsid400210ajtcl5s45nb
slug: hosting-your-serverless-website-using-aws
cover: https://cdn.hashnode.com/res/hashnode/image/upload/v1723186375254/40c205d5-4dc8-4f1d-b4cb-0e430e3b0946.png
ogImage: https://cdn.hashnode.com/res/hashnode/image/upload/v1723186550964/90849f8a-e162-4fdb-a752-1e81c93442d0.png
tags: aws, cloudfront, dynamodb, aws-lambda, aws-cloudfront, s3, route-53, lambda-function, route53, aws-s3, s3-bucket, aws-dynamodb, s3-static-website-hosting, hosting-static-website-using-aws-s3, aws-boto3-dynamodb

---

**Note**: This guide requires a web domain to host the website.

## Step 1: Creating a S3 Bucket

1. Open the **AWS Management Console** and navigate to the **S3.**
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723180942878/27d7d370-e3a8-4062-b07f-88f5951a019b.png align="center")
    
2. Click **Create Bucket**.
    
3. Give it a suitable name.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723180963058/5c5e37d2-4057-4eb5-841b-b8c1f034a179.png align="center")
    
4. Scroll and click **Create Bucket**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723180989963/15364da7-a1d3-402e-94e6-c0802aef59f8.png align="center")
    
5. Copy the following and create an **index.html file**.
    
    ```bash
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Demo Website</title>
    </head>
    <body>
        <div class="container">
            <h1>You're truly welcome!</h1>
            <h2>Do like the Blog if you find the tutorial helpful.</h2>
            <ul class="social-links">
                <li><a href="https://arishahmad.hashnode.dev/hosting-your-serverless-website-using-aws" target="_blank">Blog</a></li>
            </ul>
            <h3 id="views">Views</h3>
        </div>
        <script src="script.js"></script>
    </body>
    </html>
    ```
    
6. Upload this file in the S3 bucket.
    

## Step 2: Creating CloudFront distribution.

1. Navigate to **CloudFront** in the AWS Management Console.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181014148/ecbfc07d-a021-4991-a6af-404d7a06dfe3.png align="center")
    
2. Click **Create a CloudFront distribution**.
    
3. Select the S3 bucket you just created under the **Origin domain**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181034428/dd9b307f-c703-42be-ae33-737d9d6d4d86.png align="center")
    
4. Under **Origin access** click **Origin access control settings**.
    
5. Click **Create new OAC** -&gt; **Create**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181080552/72728f42-3464-4523-bad7-834d11e31d21.png align="center")
    
6. Under **Web Application Firewall (WAF)** select **Do not enable security protections**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181100511/13bc2eb4-e04b-4bfd-92a2-94b88fb392c3.png align="center")
    
7. Scroll and click **Create distribution**.
    
    This screen will pop up.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181210821/00889a72-1042-4d6e-9af9-f01d2d01a72b.png align="center")
    
8. Click **Copy policy**.
    
    If the **Copy policy** option doesn't appear.
    
    1. Open the distribution.
        
    2. Click on **Origins**.
        
        ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181321099/f8e5e9b9-833e-4bca-9993-f07619e03530.png align="center")
        
    3. Select the origin and click **Edit**.
        
    4. Scroll and click **Copy policy**.
        
        ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181343292/31155ce6-11d8-4ce7-9a51-ca41ae6ea174.png align="center")
        

## Step 3: Updating S3 bucket policy

1. Navigate to **S3** in the AWS Management Console.
    
2. Open the bucket.
    
3. Click on **Permissions**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181250009/b4bfb021-48b1-436f-9d14-47c5881b42ca.png align="center")
    
4. On **Bucket policy** click **Edit**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181267320/d0b4423d-8fdd-4fd5-b580-0d7aac4b57bf.png align="center")
    
5. Paste the policy and click **Save changes**.
    

## Step 4: Set up Route53

1. Navigate to **Route 53 Dashboard** in the AWS Management Console.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181399514/030fcce2-2e14-423e-a4f3-e4bb1eb76843.png align="center")
    
2. Click **Create hosted zone**.
    
3. Enter the **Domain name.**
    
4. Select **Public hosted zone** under **Type**.
    
5. Scroll and click **Create hosted zone**.
    
6. Select the record with NS type, Record details will open up.
    
7. Copy each **value** of this **record** and add this to the **Nameserver**. This option will be provided by the platform where you bought the domain.
    

## Step 5: Editing the CloudFront

1. Open the CloudFront distribution.
    
2. Select the distribution and Click **Edit** on **Settings**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181538084/dff92a15-0087-4d7a-b53e-3db2cb603b97.png align="center")
    
3. Select **Add items** under **Alternate domain name (CNAME)**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181578394/215dc090-b5bb-455e-9697-1672da8eadea.png align="center")
    
4. Enter the domain on which you want to host the website (Example: If your domain name is example.com then enter demo.example.com).
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181702881/656dd2a3-f43f-48db-98ef-33069b8a6633.png align="center")
    
5. Scroll and click **Request certificate**. A new tab will open up.
    
6. Click **Next**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181729807/dd90c575-519c-4710-983b-4459a28d577f.png align="center")
    
7. Enter **.&lt;your domain name&gt;** (Example: If your domain name is example.com, enter \*.example.com).
    
8. Scroll and click **Request**.
    
    Your new certificate might continue to display a status of **Pending validation** for **up to 30 minutes.**
    
9. Open the Certificate.
    
10. Click **Create records in Route 53**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181868451/1b3018bf-2ec8-4fe1-8930-29cb3542489a.png align="center")
    
11. Click **Create records**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723183474199/71171552-cabc-4f06-8eca-36c94f7bf951.png align="center")
    
    You can check the new route must be added in the hosted zone.
    
    Now wait till the status of certificate is not **Issued**.
    
12. Go back to the CloudFront tab, and select the newly created Custom SSL certificate.
    
13. Enter **index.html** under the **Default root object.**
    
14. Scroll and click **Save changes**.
    

## Step 6: Create a record in the Hosted Zone.

1. Open the hosted zone in Route53.
    
2. Click **Create record**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723181979984/8c6fd757-0303-4101-a34d-efd3e16e69c6.png align="center")
    
3. In record name enter the subdomain. This should be same as the alternate domain name you entered in the CloudFront distribution.
    
4. Enable the **Alias**.
    
5. Choose **Alias to CloudFront distribution** and select the distribution you created.
    
6. Click **Create Record**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182181976/99b7243f-59a1-41e2-9a39-7a50cf49275d.png align="center")
    
7. Now visit the alternate domain name, and you will see the hosted website. This might take few minutes until the CloudFront distribution is ready.
    

## Step 7: Set up the DynamoDB

1. Navigate to **DynamoDB** in the **AWS Management Console**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182313168/686f5a64-c8a2-4774-83ed-9ecc253ae034.png align="center")
    
2. Click **Create table**.
    
3. Enter a suitable name.
    
4. Endet **id** in **Partition key**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182337519/c8a3e614-86ff-47dd-be61-919dafc05221.png align="center")
    
5. Scroll and click **Create Table**.
    
6. Select the table, go to **Actions** -&gt; **Explore items**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182357452/00e32686-08ca-4b14-ad8e-9ebc585a0f62.png align="center")
    
7. Click **Create item**.
    
8. Enter value **0** in id Attribute.
    
9. Select **Add new attribute** -&gt; **Number**.
    
10. Enter the Attribute name as views and enter value **0**.
    
11. Click **Create item**.
    

## Step 8: Create the Lambda Function

1. Navigate to **AWS Lambda** in the **AWS Management Console**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182407141/9d8fec66-df7d-487c-a993-80efaabf26ff.png align="center")
    
2. Click **Create a function**.
    
3. Enter a suitable **Function name**.
    
4. Select Python 3.8 in Runtime.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182457575/e149f5ef-d6f2-4947-9421-0a071c8e0650.png align="center")
    
5. Select **Create a new role with basic Lambda permissions** in the **Execution role**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182486414/2639886d-6292-474e-9d7e-eb06895d0ef8.png align="center")
    
6. In Advanced settings click **Enable function URL** and **NONE** in **Auth Type.**
    
7. Check the **Configure cross-origin resource sharing (CORS)**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182587671/7971746b-3dfb-4315-b596-fb407f0ac71b.png align="center")
    
8. Scroll and click **the Create function**.
    
9. Copy the following code.
    
    ```bash
    import json
    import boto3
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table('demo-table')
    def lambda_handler(event, context):
        response = table.get_item(Key={
            'id':'0'
        })
        views = response['Item']['views']
        views = views + 1
        
        print(views)
        
        response = table.put_item(Item={
            'id':'0',
            'views': views
        })
        
        return views
    ```
    
10. Paste this file in the code area.
    
    **Note**: Change the table name in this file.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182623973/d02db530-6d0b-4318-8454-e8261e02d602.png align="center")
    
11. Save the file then Click **Deploy**.
    
12. Click **Configuration** -&gt; **Permissions**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182655650/d7f015b0-902c-48e3-8b97-9e04bada17dd.png align="center")
    
13. Click on the role, a new tab will open up.
    
14. Click **Addpermissions** -&gt; **Attach policies**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182702653/9d9b3bd8-f05f-4a62-8bab-5d0109369616.png align="center")
    
15. Search and select **AmazonDynamoDBFullAccess**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182732723/82b65160-768c-4e04-8077-5a3bc3ffe2e2.png align="center")
    
16. Click **Add permissions**.
    
17. Now Come back to the Lambda tab.
    
18. Click Code and then click **Test**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723182915105/59e3b824-23d8-4e15-b4a0-12dbf6bc3376.png align="center")
    
19. Enter a suitable name, then click **Save**.
    
20. Now click **Test** again, and you can see the response.
    

## Step 9: Creating the JS file

1. Copy the following code.
    
    ```bash
    let counter = document.getElementById("views");
    
    async function updateCounter() {
        try {
            let response = await fetch("lambda_function_url");
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            let data = await response.json();
            console.log(data);
            counter.innerHTML = `Views: ${data}`;
        } catch (error) {
            console.error('Failed to fetch data:', error);
            counter.innerHTML = 'Failed to load views';
        }
    }
    
    updateCounter();
    ```
    
2. Create a file named script.js file, and paste the code.
    
3. Copy the Lambda function URL and replace lambda\_function\_url in the file with the URL.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723183074693/cc0d2f6e-8e61-4183-acbc-7022c562263c.png align="center")
    
4. Now upload this script.js file in the S3 bucket.
    

## Step 10: Finishing Up

Visit the Alternate domain name you entered in the CloudFront. You can see the hosted website and with each reload the view counter will increase. You can also check the increased value of views in the dynamoDB.

If the changes are not visible you may need to create an **Invalidation.**

1. Open the CloudFront distribution you created.
    
2. Click **Invalidations** -&gt; **Create invalidation**.
    
    ![](https://cdn.hashnode.com/res/hashnode/image/upload/v1723183110249/cde7d336-e98b-4f13-8366-5f823298bc28.png align="center")
    
3. Type /\* and click **Create invalidation**.
    
4. Wait while the status is in progress.
    
5. Reload the website, and you'll be able to see the changes.
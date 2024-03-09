# Welcome to the "Donations to Orphanages" project
This project is a backend server written in Go for a web application that allows users to find information about the needs of orphanages and make donations.

## Functionality
- Register and sign in
- View orphanage's needs
- Search orphanages by region and type of need filter
- Donate money to orphanages
- Subscribe to monthly donations
- Read and write commentaries
- Administrators can add, modify, and delete orphanage needs.

## Requirements
- Go 1.15 or higher
- MongoDB

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/olzhjans/Donation-App
   ```
2. Install GO from https://go.dev/dl/
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Install Mongo Database Server. Instructions: https://www.mongodb.com/docs/v6.0/installation/
5. Install Mongo Database Tools. Instructions: https://www.mongodb.com/docs/database-tools/installation/installation/
6. Restore database:
   ```bash
   mongorestore --host <your-host> --port <your-port> <path-to-backup-copy>
   ```
7. Run the server:
   ```bash
   go run main.go
   ```

## Using the API

### Signing in
`POST /login`
Example request body:
```json
{
    "phone":"+77787787878",
    "password":"12345678"
}
```
- `phone`: The phone number (login) of user.
- `password`: The password.

### User signing up
`POST /userSignUp`
Example request body:
```json
{
    "phone":"+77766787878",
    "password":"12345678",
    "firstname":"John",
    "lastname":"Doe",
    "email":"johndoe@gmail.com",
    "region":"Almaty"
}
```

### Admin signing up
`POST /adminSignUp`
Example request body:
```json
{
   "phone":"+77766787878",
   "password":"12345678",
   "firstname":"John",
   "lastname":"Doe",
   "email":"johndoe@gmail.com",
   "region":"Almaty",
   "id":"950101123123",
   "who":"Moderator",
   "orphanage-id":"65ba7800ab0265f4fa9d4b60"
}
```
- `id`: The ID of Kazakhstan citizen.
- `who`: Is he `Admin` or `Moderator`.
- `orphanage-id`: If he is `Moderator`, then orphanage's ID is required.
**Note**: these data will be added to `waitinglist` collection.

### Waiting list review
`GET /showWaitingList`
#### `waitinglist` collection will be shown

### Confirming registration from waiting list by `_id`
`GET /confirmRegistration?_id=ENTER_ID`
#### Registration will be confirmed. Email will be sent.
**Note**: enter actual `_id`.

### Rejecting registration from waiting list by `_id`
`DEL /deleteWaitingList?_id=ENTER_ID`
#### Registration will be rejected. Email will be sent.

### Editing user data
`POST /editUser`
```json
{
    "_id":"65e719fd9c39339929ae5b5d",
    "phone":"+77766787878",
    "password":"12345678",
    "firstname":"John",
    "lastname":"Doe",
    "email":"johndoe@gmail.com",
    "region":"Almaty",
    "donated":0
}
```
**Note**: you can only enter the fields you want to change, but `"_id"` is required.

### Editing admin data
`POST /editAdmin`
```json
{
   "_id":"65e71c90cbb401656adb55a6",
   "phone":"+77766787878",
   "password":"12345678",
   "firstname":"John",
   "lastname":"Doe",
   "email":"johndoe@gmail.com",
   "region":"Almaty",
   "id":"950101123123",
   "who":"Admin",
   "orphanage-id":"65ba7800ab0265f4fa9d4b60"
}
```
**Note**: you can only enter the fields you want to change, but `"_id"` is required.

### Adding new orphanage
`POST /addOrphanage`
```json
{
   "name":"Umit",
   "region":"Almaty",
   "address":"Abay st. 1",
   "description":"Description of orphanage",
   "childs-count":40,
   "working-hours":"8AM - 6PM",
   "photos":["link_to_photo_1","link_to_photo_2"],
   "bill":0
}
```

### Editing an orphanage
`POST /editOrphanage`
```json
{
    "_id":"65ebfa792a10e66ce0a6a8f5",
   "name":"Umit",
   "region":"Almaty",
   "address":"Abay st. 2",
   "description":"Description of orphanage",
   "childs-count":50,
   "working-hours":"8AM - 6PM",
   "photos":["link_to_photo_1","link_to_photo_2"],
   "bill":0
}
```
**Note**: `"_id"` is required.

### Search orphanage by name
`GET /getOrphanage?name=ENTER_NAME`

### Search needs by filter (region and type)
`POST /getNeedsByRegionAndType`
```json
{
    "region":"Almaty",
    "category-of-donate":"Clothes"
}
```

### Show `wherespent` data by orphanage's id and time filter
`POST /showWhereSpent`
```json
{
    "orphanage-id":"65c07ee1dfda391f6fe449a6",
    "from":"2024-01-23T00:00:00Z",
    "to":"2024-02-24T00:00:00Z"
}
```

### Show needs by orphanage `"_id"`
`GET /showNeeds?orphanageid=ENTER_ID`

### Add new need
`POST /addNeed`
```json
{
"amount":"1",
"categoryofdonate":"1",
"sizeofclothes":"1",
"typeofcount":"1",
"typeofdonate":"1",
"orphanageid":"1",
"isactive":true
}
```

### Activating need
`GET /activateNeed?needid=ENTER_ID`

### Deactivating need
`GET /deactivateNeed?needid=ENTER_ID`

### Show comments by need's `"_id"` and time filter
`POST /getComments`
```json
{
    "need-id":"65ba75f0ab0265f4fa9d4b5f",
    "from":"2024-02-20T00:00:00Z",
    "to":"2024-04-23T00:00:00Z"
}
```

### Add commentary
`POST /addComment`
```json
{
    "need-id":"65c0842fdfda391f6fe449ac",
    "user-id":"65c3b0b072629755858ece76",
    "text":"text"
}
```

### Delete commentary
`DEL /deleteComment?_id=ENTER_COMMENTARY_ID`

### Donation
`POST /addDonate`
```json
{
    "bankdetails-id":"65bbbe64a29af2768a9009cb",
    "orphanage-id":["65ba7800ab0265f4fa9d4b60","65c07ee1dfda391f6fe449a6"],
    "sum":15000
}
```
**Note**: sum of donate will be charged off from user's bill, also will be added to orphanage's bill. Donation data will be added to `donationhistory` collection.

### Show summary donated to orphanage
`POST /getTotalDonatedByOrphanageIdAndPeriod`
```json
{
    "id":"65c07ee1dfda391f6fe449a6",
    "from":"2024-01-01T00:00:00Z",
    "to":"2024-12-31T00:00:00Z"
}
```

### Show summary donated by user
`POST /getTotalDonatedByUserIdAndPeriod`
```json
{
    "id":"65c1da85b9683cf7113767cf",
    "from":"2024-01-01T00:00:00Z",
    "to":"2024-12-31T00:00:00Z"
}
```

### Add donation subscribe
`POST /addDonationSubscribe`
```json
{
    "bank-details":{
        "name":"FIRSTNAME LASTNAME",
        "expiring":"01.28",
        "cvv":"777",
        "cardnumber":"1234123412341234",
        "userid":"65c1da85b9683cf7113767cf",
        "bill":5000
    },
    "orphanageid":["65ba7800ab0265f4fa9d4b60"],
    "amount":5000,
    "whichday":20,
    "isactive":true
}
```
**Note**: bank card will be saved to `bankdetails` collection.

### Deactivating donate subscription
`GET /deactivateDonateSubscription?_id=ENTER_ID`

### Show donation subscribes by user
`GET /getDonationSubscribeByUserId?userid=ENTER_ID`

### Show donation history by user
`GET /getDonationHistoryByUserId`

# TP FINAL

## INFORMATIONS
Go version 1.23.2

## HOW TO RUN
- git clone <URL_DU_DEPOT>
- cd <NOM_DU_DEPOT>
- go mod tidy
- go run .

## Routes

### CATS

#### Create a Cat
- URL: /cats
- Method: POST
- Description: Create a new cat.
- Request Body:
```json
{
  "name": "Legroschat",
  "age": 11,
  "race": "onsaitpo",
  "gender": "male",
  "weight": 8.7
}
```
- Response:
```json
{
  "id": 1,
  "name": "Legroschat",
  "age": 11,
  "race": "onsaitpo",
  "gender": "male",
  "weight": 8.7
}
```

#### Get All Cats
- URL: /cats
- Method: GET
- Description: Retrieve all cats.
- Response:
```json
[
  {
    "id": 1,
    "name": "Legroschat",
    "age": 11,
    "race": "onsaitpo",
    "gender": "male",
    "weight": 8.7
  }
]
```

#### Get Cat by ID
- URL: /cats/{id}
- Method: GET
- Description: Retrieve a cat by its ID.
- Response:
```json
{
  "id": 1,
  "name": "Legroschat",
  "age": 11,
  "race": "onsaitpo",
  "gender": "male",
  "weight": 8.7
}
```

#### Update a Cat
- URL: /cats/{id}
- Method: PUT
- Description: Update a cat.
- Request Body:
```json
{
  "name": "Lemoyenchat",
  "age": 11,
  "race": "onsaitpo",
  "gender": "male",
  "weight": 8.7
}
```
- Response:
```json
{
  "id": 1,
  "name": "Lemoyenchat",
  "age": 11,
  "race": "onsaitpo",
  "gender": "male",
  "weight": 8.7
}
```

#### Delete a Cat
- URL: /cats/{id}
- Method: DELETE
- Description: Delete a cat.
- Response:
```json
{
  "message": "Cat deleted"
}
```

#### Get Cat History
- URL: /cats/{id}/history
- Method: GET
- Description: Retrieve the visit history of a cat.
- Response:
```json
[
  {
    "id": 1,
    "cat_id": 1,
    "date": "2023-10-01",
    "reason": "Routine check-up",
    "doctor": "Dr. Smith",
    "treatments": [
      {
        "id": 1,
        "visit_id": 1,
        "medicine": "Antibiotic",
        "start_date": "2023-10-01",
        "end_date": "2023-10-10",
        "doctor_tip": "Take with food"
      }
    ]
  }
]
```

#### Get Visits by Cat
- URL: /cats/{id}/visits
- Method: GET
- Description: Retrieve all visits for a specific cat.
- Response:
```json
[
  {
    "id": 1,
    "cat_id": 1,
    "date": "2023-10-01",
    "reason": "Routine check-up",
    "doctor": "Dr. Smith"
  }
]
```

#### Get Visit by Cat
- URL: /cats/{id}/visits/{visitId}
- Method: GET
- Description: Retrieve a specific visit for a cat.
- Response:
```json
{
  "id": 1,
  "cat_id": 1,
  "date": "2023-10-01",
  "reason": "Routine check-up",
  "doctor": "Dr. Smith"
}
```

#### Get Treatments by Cat by Visit
- URL: /cats/{id}/visits/{visitId}/treatments
- Method: GET
- Description: Retrieve all treatments for a specific visit of a cat.
- Response:
```json
[
  {
    "id": 1,
    "visit_id": 1,
    "medicine": "Antibiotic",
    "start_date": "2023-10-01",
    "end_date": "2023-10-10",
    "doctor_tip": "Take with food"
  }
]
```

#### Get Treatment by Cat by Visit
- URL: /cats/{id}/visits/{visitId}/treatments/{treatmentId}
- Method: GET
- Description: Retrieve a specific treatment for a visit of a cat.
- Response:
```json
{
  "id": 1,
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```

### VISITS

#### Create a Visit
- URL: /visits
- Method: POST
- Description: Create a new visit.
- Request Body:
```json
{
  "cat_id": 1,
  "date": "2023-10-01",
  "reason": "Routine check-up",
  "doctor": "Dr. Smith"
}
```
- Response:
```json
{
  "id": 1,
  "cat_id": 1,
  "date": "2023-10-01",
  "reason": "Routine check-up",
  "doctor": "Dr. Smith"
}
```

#### Get All Visits
- URL: /visits
- Method: GET
- Description: Retrieve all visits.
- Response:
```json
[
  {
    "id": 1,
    "cat_id": 1,
    "date": "2023-10-01",
    "reason": "Routine check-up",
    "doctor": "Dr. Smith"
  }
]
```

#### Get Visit by ID
- URL: /visits/{id}
- Method: GET
- Description: Retrieve a visit by its ID.
- Response:
```json
{
  "id": 1,
  "cat_id": 1,
  "date": "2023-10-01",
  "reason": "Routine check-up",
  "doctor": "Dr. Smith"
}
```

#### Update a Visit
- URL: /visits/{id}
- Method: PUT
- Description: Update a visit.
- Request Body:
```json
{
  "cat_id": 1,
  "date": "2023-10-02",
  "reason": "Follow-up",
  "doctor": "Dr. Smith"
}
```
- Response:
```json
{
  "id": 1,
  "cat_id": 1,
  "date": "2023-10-02",
  "reason": "Follow-up",
  "doctor": "Dr. Smith"
}
```

#### Delete a Visit
- URL: /visits/{id}
- Method: DELETE
- Description: Delete a visit.
- Response:
```json
{
  "message": "Visit deleted"
}
```

#### Get Treatments by Visit
- URL: /visits/{id}/treatments
- Method: GET
- Description: Retrieve all treatments for a specific visit.
- Response:
```json
[
  {
    "id": 1,
    "visit_id": 1,
    "medicine": "Antibiotic",
    "start_date": "2023-10-01",
    "end_date": "2023-10-10",
    "doctor_tip": "Take with food"
  }
]
```

#### Get Treatment by Visit
- URL: /visits/{id}/treatments/{treatmentId}
- Method: GET
- Description: Retrieve a specific treatment for a visit.
- Response:
```json
{
  "id": 1,
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```

### TREATMENTS

#### Create a Treatment
- URL: /treatments
- Method: POST
- Description: Create a new treatment.
- Request Body:
```json
{
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```
- Response:
```json
{
  "id": 1,
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```

#### Get All Treatments
- URL: /treatments
- Method: GET
- Description: Retrieve all treatments.
- Response:
```json
[
  {
    "id": 1,
    "visit_id": 1,
    "medicine": "Antibiotic",
    "start_date": "2023-10-01",
    "end_date": "2023-10-10",
    "doctor_tip": "Take with food"
  }
]
```

#### Get Treatment by ID
- URL: /treatments/{id}
- Method: GET
- Description: Retrieve a treatment by its ID.
- Response:
```json
{
  "id": 1,
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```

#### Update a Treatment
- URL: /treatments/{id}
- Method: PUT
- Description: Update a treatment.
- Request Body:
```json
{
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```
- Response:
```json
{
  "id": 1,
  "visit_id": 1,
  "medicine": "Antibiotic",
  "start_date": "2023-10-01",
  "end_date": "2023-10-10",
  "doctor_tip": "Take with food"
}
```

#### Delete a Treatment
- URL: /treatments/{id}
- Method: DELETE
- Description: Delete a treatment.
- Response:
```json
{
  "message": "Treatment deleted"
}
```
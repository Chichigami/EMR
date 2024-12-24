# Frontend

## Login

- Login
  - Username
  - Password
  - Optional Clinic/Server
- Logo

## Dashboard

- Show who is logged in
  - A way to quick swap login
- State of the current day
  - Time and name of patient
  - "Rooms" that patient are in
- Interactive calendar to change current day. To check past and present dates.
- Patient search via query
  - ID, DoB, Name

## Patient Info

- GET request for info + first chart
- Show special note
- Image of patient
- ID
- Last, First name
- DoB
- Insurance
- Primary Care Physician (PCP)
- Pharmacy

## Charting

- Load latest chart (should be last visit)

  - GET req to fetch old charts

- Button to make new chart
- Symptoms
- Diagnoses
  - Active vs Inactive problems. Show at top of chart. Button. Bool
- Medication
  - Load medication from previous chart recommendations
- Billing
  - ICD-10 codes
    - Cache all codes if ICD 10 codes if possible. query string on the fly if possible
- Signature of current user
- On exit: POST request to backend

- Special note like VIP?
- Allergies?

# Backend

## Process

1. Login
2. Dashboard
3. Patient info
4. Charting
5. Upload attachments

## Charting

- handlerFetchPrevious: from patient info
- handlerFetchAll: from charting

## PostgreSQL Database

## Queries

- INSERT login
- UPDATE login
- DELETE login

- Insert patient
- SELECT patient
- Update patient
- DELETE patient

- Insert chart
- SELECT chart
- UPDATE chart
- DELETE chart

## Dashboard

- Fetch all patients

### Patient demogrpahic

- ID UUID
- First name
- Last name
- DoB
- SSN?
- Insurance

### Doctor login

- New login
  - Username
  - Password -> hashed -> force high entropy, maybe some physical key option?
- Login verification
  - Password hash
- Authorization tiers
  - Doctor vs Nurse vs Technician

# Data standard?

- HL7, FHIR

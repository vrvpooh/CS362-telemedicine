# Telemedicine System Design Document

## Step 1: Data Design

### 1. Entities และ Attributes (ข้อมูลที่จำเป็น)

### User
เก็บข้อมูลพื้นฐานของผู้ใช้งาน

| Attribute | Description |
|---|---|
| userId (PK) | รหัสผู้ใช้งาน |
| username | ชื่อบัญชีผู้ใช้ |
| password | รหัสผ่าน |
| fullName | ชื่อ-นามสกุล |
| role | ประเภทผู้ใช้ (Patient / Doctor / Admin) |

---

### Patient
ข้อมูลเพิ่มเติมเฉพาะของคนไข้

| Attribute | Description |
|---|---|
| patientId (PK/FK) | รหัสคนไข้ |
| bloodGroup | กรุ๊ปเลือด |
| allergies | ประวัติการแพ้ยา/อาหาร |
| medicalHistory | ประวัติการรักษา |

---

### Doctor
ข้อมูลความเชี่ยวชาญของแพทย์

| Attribute | Description |
|---|---|
| doctorId (PK/FK) | รหัสแพทย์ |
| specialty | ความเชี่ยวชาญ |
| education | ประวัติการศึกษา |
| rating | คะแนนรีวิว |
| isVerified | สถานะการยืนยันตัวตน |

---

### TimeSlot
ช่วงเวลาที่แพทย์เปิดให้จอง

| Attribute | Description |
|---|---|
| slotId (PK) | รหัสช่วงเวลา |
| doctorId (FK) | รหัสแพทย์ |
| startTime | เวลาเริ่มต้น |
| endTime | เวลาสิ้นสุด |
| status | สถานะ (Available / Booked) |

---

### Appointment
ข้อมูลการนัดหมายและการเชื่อมต่อ Zoom

| Attribute | Description |
|---|---|
| appId (PK) | รหัสการนัดหมาย |
| patientId (FK) | รหัสคนไข้ |
| doctorId (FK) | รหัสแพทย์ |
| slotId (FK) | รหัสช่วงเวลา |
| status | สถานะ (Pending / Confirmed / Cancelled) |
| zoomLink | ลิงก์สำหรับการประชุม |

---

### MedicalRecord
บันทึกผลการตรวจหลังจบการสนทนา

| Attribute | Description |
|---|---|
| record_id (PK) | รหัสบันทึกการรักษา |
| appId (FK) | รหัสการนัดหมาย |
| diagnosis | การวินิจฉัยโรค |
| prescription | ใบสั่งยา |
| doctorNote | หมายเหตุจากแพทย์ |

---

### 2. Relationships (ความสัมพันธ์ระหว่างข้อมูล)

- **User ↔ Patient / Doctor**  
  เป็นความสัมพันธ์แบบ **Inheritance (One-to-One)**  
  - User 1 คน จะเป็นได้ทั้ง Patient หรือ Doctor

- **Doctor ↔ TimeSlot**  
  ความสัมพันธ์แบบ **One-to-Many**  
  - แพทย์ 1 คน สามารถสร้างช่วงเวลาว่างได้หลายช่วง

- **Patient ↔ Appointment**  
  ความสัมพันธ์แบบ **One-to-Many**  
  - คนไข้ 1 คน สามารถมีการนัดหมายได้หลายครั้ง

- **Doctor ↔ Appointment**  
  ความสัมพันธ์แบบ **One-to-Many**  
  - แพทย์ 1 คน สามารถรับนัดหมายจากหลายคนไข้

- **TimeSlot ↔ Appointment**  
  ความสัมพันธ์แบบ **One-to-One**  
  - 1 ช่วงเวลาที่ถูกจอง จะสัมพันธ์กับ 1 การนัดหมายเท่านั้น  
  - เพื่อป้องกันการจองซ้อน

- **Appointment ↔ MedicalRecord**  
  ความสัมพันธ์แบบ **One-to-One**  
  - การนัดหมาย 1 ครั้งที่เสร็จสิ้น จะมีบันทึกการรักษาได้ 1 ฉบับ

## Step 2 Architecture Diagram

### Layered Architecture
* **Controller Layer**
    * `AuthController`
    * `DoctorController`
    * `AppointmentController`
    * `RecordController`
* **Service Layer**
    * `AuthService`
    * `DoctorService`
    * `AppointmentService`
    * `ZoomService`
    * `MedicalRecordService`
* **Repository Layer**
    * `UserRepository`
    * `DoctorDetailRepository`
    * `TimeSlotRepository`
    * `AppointmentRepository`
    * `MedicalRecordRepository`

---

### Folder Structure Example
```text
telemedicine-system/
├── controller/
│   ├── AuthController.java
│   ├── DoctorController.java
│   ├── SlotController.java 
│   ├── AppointmentController.java
│   ├── RecordController.java      
│   └── NotificationController.java
├── service/
│   ├── AuthService.java
│   ├── AppointmentService.java 
│   └── ZoomIntegrationService.java
├── repository/
│   ├── UserRepository.java
│   ├── DoctorRepository.java
│   ├── SlotRepository.java
│   └── AppointmentRepository.java
├── model/
│   ├── User.java
│   ├── Appointment.java
│   └── MedicalRecord.java
└── main application
```
---

## Step 3 Contract

### Request Contract (POST `/api/appointments`)

```json
{
  "patientId": "P101",
  "doctorId": "D202",
  "slotId": "S999",
  "symptoms": "ปวดหัวต่อเนื่อง 3 วัน"
}
```
Java
```
public interface AppointmentService {
    AppointmentResponse createAppointment(String patientId, String slotId);
    void updateStatus(String appointmentId, String status);
}
```
Expected Response
```
{
  "appointmentId": "APP-555",
  "status": "CONFIRMED",
  "zoomToken": "eyJhbGciOiJIUzI1NiJ9...", 
  "appointmentTime": "2026-03-20 09:00",
  "message": "ระบบออกตั๋วและเตรียมห้อง Video Call เรียบร้อยแล้ว"
}
```

## Table Design
### Users table
| Field | Type | Description |
| :--- | :--- | :--- |
| **user_id** | UUID (PK) | ไอดีหลักของผู้ใช้ |
| **username** | String | ชื่อผู้ใช้งานสำหรับ Login |
| **password_hash** | String | รหัสผ่านที่เข้ารหัสแล้ว |
| **role** | Enum | บทบาท: `PATIENT`, `DOCTOR`, `ADMIN` |
| **full_name** | String | ชื่อ-นามสกุล |
| **medical_info** | JSON/Text | ข้อมูลแพ้อาหาร หรือกรุ๊ปเลือด |

### Doctors table
| Field | Type | Description |
| :--- | :--- | :--- |
| **doctor_id** | UUID (PK, FK) | เชื่อมกับ `Users.user_id` |
| **specialty** | String | ความเชี่ยวชาญ เช่น อายุรกรรม, จิตเวช |
| **biography** | Text | ประวัติการศึกษาและการทำงาน |
| **is_verified** | Boolean | สถานะการตรวจสอบใบประกอบวิชาชีพโดย Admin |
| **rating_avg** | Float | คะแนนเฉลี่ยจากการรีวิว |

### Time_Slots table
| Field | Type | Description |
| :--- | :--- | :--- |
| **slot_id** | UUID (PK) | ไอดีช่วงเวลา |
| **doctor_id** | UUID (FK) | ไอดีแพทย์เจ้าของตาราง |
| **start_time** | DateTime | เวลาเริ่มต้น (เช่น 09:00) |
| **end_time** | DateTime | เวลาสิ้นสุด (เช่น 10:00) |
| **is_available** | Boolean | สถานะว่าง/ไม่ว่าง เพื่อป้องกันการจองซ้ำ |

### Appointments table
| Field | Type | Description |
| :--- | :--- | :--- |
| **app_id** | UUID (PK) | ไอดีการนัดหมาย |
| **patient_id** | UUID (FK) | ไอดีคนไข้ผู้จอง |
| **doctor_id** | UUID (FK) | ไอดีแพทย์ที่ถูกจอง |
| **slot_id** | UUID (FK) | ไอดีช่วงเวลาที่เลือก |
| **status** | Enum | สถานะ: `PENDING`, `CONFIRMED`, `COMPLETED`, `CANCELLED` |
| **zoom_link** | String | ลิงก์สำหรับเข้าห้อง Video Call |

### Medical_Records table
| Field | Type | Description |
| :--- | :--- | :--- |
| **record_id** | UUID (PK) | ไอดีบันทึกการรักษา |
| **app_id** | UUID (FK) | เชื่อมกับการนัดหมายครั้งนั้นๆ |
| **diagnosis** | Text | สรุปอาการและคำแนะนำจากแพทย์ |
| **prescription** | Text | ข้อมูลใบสั่งยาเบื้องต้น |
| **created_at** | DateTime | วันที่บันทึกข้อมูล |

### Feedbacks table
| Field | Type | Description |
| :--- | :--- | :--- |
| **feedback_id** | UUID (PK) | ไอดีการรีวิว |
| **app_id** | UUID (FK) | เชื่อมกับการนัดหมายที่จบไปแล้ว |
| **rating** | Integer | คะแนน (1-5) |
| **comment** | Text | ความคิดเห็นจากคนไข้ |

## Refinement

### 1. Entity: User / Patient / Doctor
เน้นเรื่องการตรวจสอบความถูกต้องของข้อมูลพื้นฐาน

### Methods

| Method | Description |
|---|---|
| isValidPassword() | ตรวจสอบว่ารหัสผ่านมีความยาวและอักขระครบตามนโยบายความปลอดภัย |
| getFullName() | รวมชื่อและนามสกุลเพื่อนำไปแสดงผลบนหน้า UI |
| isMedicalVerified() | (สำหรับ Doctor) ตรวจสอบว่าใบประกอบวิชาชีพผ่านการอนุมัติหรือยัง |

---

### 2. Entity: TimeSlot
เน้นเรื่องการจัดการช่วงเวลาว่างของแพทย์

### Methods

| Method | Description |
|---|---|
| isExpired() | ตรวจสอบว่าช่วงเวลานี้ผ่านเลยเวลาปัจจุบันไปแล้วหรือยัง (ถ้าผ่านแล้วจะไม่แสดงให้จอง) |
| calculateDuration() | คำนวณระยะเวลาของ Slot เช่น 30 นาที หรือ 1 ชั่วโมง |
| updateStatus(newStatus) | เปลี่ยนสถานะของ Slot เช่น จาก `Available` เป็น `Reserved` |

---

### 3. Entity: Appointment
เน้นเรื่องการจัดการสถานะการนัดหมายและการเชื่อมต่อระบบประชุม

### Methods

| Method | Description |
|---|---|
| canCancel() | ตรวจสอบว่าคนไข้ยังสามารถยกเลิกนัดได้หรือไม่ (เช่น ต้องยกเลิกก่อนเวลาเริ่มอย่างน้อย 24 ชั่วโมง) |
| isReadyForConsultation() | ตรวจสอบว่าถึงเวลาเข้าห้องประชุมหรือยัง (เช่น เปิดให้เข้าก่อน 5–10 นาที) |
| updateAppointmentStatus(status) | อัปเดตสถานะ เช่น `Confirmed`, `Completed`, หรือ `Cancelled` |

---

### 4. Entity: MedicalRecord / Feedback
เน้นการจัดการข้อมูลหลังการรักษา

### Methods

| Method | Description |
|---|---|
| hasPrescription() | ตรวจสอบว่าในบันทึกนี้มีใบสั่งยาแนบมาด้วยหรือไม่ |
| isValidRating() | ตรวจสอบว่าคะแนน Feedback อยู่ในช่วง 1–5 ดาว |

---

## Example: Class Diagram (Refinement)

```plaintext
Appointment
--------------------------------
- appId: String
- status: Enum
- startTime: DateTime
--------------------------------
+ canCancel(): Boolean
+ isReadyForConsultation(): Boolean
+ updateStatus(status: String)

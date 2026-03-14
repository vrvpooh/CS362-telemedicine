Task 1: Step 2 Architecture Diagram
Layered Architecture
Controller Layer
 └── AuthController
 └── DoctorController
 └── AppointmentController
 └── RecordController
Service Layer
 └── AuthService
 └── DoctorService
 └── AppointmentService
 └── ZoomService
 └── MedicalRecordService
Repository Layer
 └── UserRepository
 └── DoctorDetailRepository
 └── TimeSlotRepository
 └── AppointmentRepository
 └── MedicalRecordRepository


 Folder Structure Example
 telemedicine-system/
├── controller/
│     ├── AuthController.java
│     ├── DoctorController.java
│     ├── SlotController.java 
│     ├── AppointmentController.java
│     ├── RecordController.java     
│     └── NotificationController.java
├── service/
│     ├── AuthService.java
│     ├── AppointmentService.java 
│     └── ZoomIntegrationService.java
├── repository/
│     ├── UserRepository.java
│     ├── DoctorRepository.java
│     ├── SlotRepository.java
│     └── AppointmentRepository.java
├── model/
│     ├── User.java
│     ├── Appointment.java
│     └── MedicalRecord.java
└── main application


Task 2: Step 3 Contract
Request Contract (POST /api/appointments)
{
  "patientId": "P101",
  "doctorId": "D202",
  "slotId": "S999",
  "symptoms": "ปวดหัวต่อเนื่อง 3 วัน"
}
Java
public interface AppointmentService {
    AppointmentResponse createAppointment(String patientId, String slotId);
    void updateStatus(String appointmentId, String status);
}

Expected Response
{
  "appointmentId": "APP-555",
  "status": "CONFIRMED",
  "zoomToken": "eyJhbGciOiJIUzI1NiJ9...", 
  "appointmentTime": "2026-03-20 09:00",
  "message": "ระบบออกตั๋วและเตรียมห้อง Video Call เรียบร้อยแล้ว"
}

Table Design
Users table
| user_id | UUID (PK) | ไอดีหลักของผู้ใช้ |
| username | String | ชื่อผู้ใช้งานสำหรับ Login  |
| password_hash  String | รหัสผ่านที่เข้ารหัสแล้ว |
| role | Enum | บทบาท: PATIENT, DOCTOR, ADMIN  |
| full_name | String | ชื่อ-นามสกุล |
| medical_info | JSON/Text | ข้อมูลแพ้อาหาร หรือกรุ๊ปเลือด  |

Doctors table
| doctor_id | UUID (PK, FK) | เชื่อมกับ Users.user_id |
| specialty | String | ความเชี่ยวชาญ เช่น อายุรกรรม, จิตเวช  |
| biography | Text | ประวัติการศึกษาและการทำงาน  |
| is_verified | Boolean | สถานะการตรวจสอบใบประกอบวิชาชีพโดย Admin  |
| rating_avg | Float | คะแนนเฉลี่ยจากการรีวิว  |

Time_Slots table
| slot_id | UUID (PK) | ไอดีช่วงเวลา |
| doctor_id | UUID (FK) | ไอดีแพทย์เจ้าของตาราง  |
| start_time | DateTime | เวลาเริ่มต้น (เช่น 09:00)  |
| end_time | DateTime | เวลาสิ้นสุด (เช่น 10:00)  |
| is_available | Boolean | สถานะว่าง/ไม่ว่าง เพื่อป้องกันการจองซ้ำ  |

Appointments table
| app_id | UUID (PK) | ไอดีการนัดหมาย |
| patient_id | UUID (FK) | ไอดีคนไข้ผู้จอง  |
| doctor_id | UUID (FK) | ไอดีแพทย์ที่ถูกจอง  |
| slot_id | UUID (FK) | ไอดีช่วงเวลาที่เลือก  |
| status | Enum | สถานะ: PENDING, CONFIRMED, COMPLETED, CANCELLED  |
| zoom_link | String | ลิงก์สำหรับเข้าห้อง Video Call ที่ Generate จากระบบ  |

Medical_Records table
| record_id | UUID (PK) | ไอดีบันทึกการรักษา |
| app_id | UUID (FK) | เชื่อมกับการนัดหมายครั้งนั้นๆ  |
| diagnosis | Text | สรุปอาการและคำแนะนำจากแพทย์  |
| prescription | Text | ข้อมูลใบสั่งยาเบื้องต้น  |
| created_at | DateTime | วันที่บันทึกข้อมูล  |

Feedbacks table
| feedback_id | UUID (PK) | ไอดีการรีวิว |
| app_id | UUID (FK) | เชื่อมกับการนัดหมายที่จบไปแล้ว |
| rating | Integer | คะแนน (1-5)  |
| comment | Text | ความคิดเห็นจากคนไข้  |

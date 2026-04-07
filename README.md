# Main Structure

```sh
Farming-Simulator/
│
├── apps/
│   ├── core-api/        # Golang (main backend)
│   ├── integration/     # Node.js (WA bot, external service)
│   └── web/             # React frontend
│
├── deployments/         # docker, nginx, infra
├── scripts/             # automation
│
├── .env.example
├── .gitignore
└── README.md
```
# DB Design
---

# 🗄️ DETAIL DATABASE

## 1. users

```sql
id (uuid)
name
email
password
role (admin, partner)
created_at
```

---

## 2. farms (Lahan utama)

```sql
id (uuid)
name
location
owner_id (FK users)
created_at
```

👉 Walaupun sekarang 1 lahan, tetap buat tabel ini
→ supaya nanti bisa scale ke banyak lahan

---

## 3. sub_places (inti dari sistem kamu 🔥)

Ini yang membedakan sistem kamu

```sql
id (uuid)
farm_id (FK)
name (Kolam A, Bed 1, dll)
type (pond, plant_bed, greenhouse)
size (optional)
status (active, inactive)
created_at
```

👉 Contoh:

* Kolam Lele A
* Kolam Lele B
* Bed Cabai 1

---

## 4. cycles (Siklus budidaya)

Karena tiap sub_place bisa dipakai berulang

```sql
id (uuid)
sub_place_id (FK)
commodity_type (fish, plant)
commodity_name (lele, cabai, dll)
start_date
end_date (nullable)
status (ongoing, finished, failed)
created_at
```

👉 Ini penting banget:

* 1 kolam → bisa banyak siklus
* histori tetap tersimpan

---

## 5. activities (Log aktivitas harian)

```sql
id (uuid)
cycle_id (FK)
type (feeding, fertilizing, cleaning, note)
description
quantity (optional)
unit (kg, liter, dll)
created_by (FK users)
created_at
```

👉 Contoh:

* Pakan 5kg
* Pupuk 2 liter
* Catatan: daun menguning

---

## 6. sensor_data (IoT ready 🚀)

```sql
id (uuid)
sub_place_id (FK)
device_id
temperature
ph
humidity
created_at
```

👉 Jangan taruh di cycles:
karena sensor itu **real-time & kontinu**

---

## 7. harvests (Panen)

```sql
id (uuid)
cycle_id (FK)
harvest_date
quantity
unit (kg)
price_per_unit
total_amount
created_at
```

👉 Bisa multi panen dalam 1 siklus

---

## 8. transactions (Keuangan)

```sql
id (uuid)
cycle_id (FK)
type (expense, income)
category (pakan, pupuk, jual, dll)
amount
description
created_at
```

---

## 9. profit_sharing (Bagi hasil)

```sql
id (uuid)
cycle_id (FK)
total_profit
investor_share
operator_share
percentage_investor
percentage_operator
calculated_at
```
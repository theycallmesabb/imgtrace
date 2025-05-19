# 🖼️ imgtrace — Image Duplicate Detection App

**imgtrace** is a web application that helps you upload images and checks if the image already exists in the database using hash-based comparison. If the image is already present, it returns a match and the existing image URL; otherwise, it saves the new image.

---

## 🚀 Features

- Upload any image.
- Detects and blocks duplicate images.
- Stores unique image data in PostgreSQL.
- Uses image hashing to check similarity.
- Built with Go (Gin), GORM, and PostgreSQL.

---

## 🛠️ Tech Stack

- **Backend**: Go + Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Hashing**: pHash (Perceptual Hash)
- **File Storage**: Local folder (`static/uploads`)

---

## 🧩 How It Works

1. 📤 User uploads an image.
2. 🧠 The app generates a perceptual hash of the image.
3. 🧮 It compares the hash with existing image hashes in the database.
4. ✅ If similar image exists (hash distance ≤ 5), it returns the existing image.
5. 🆕 If not found, it saves the image and its hash to the DB.

---

## 📦 Project Structure

```bash
imgtrace/
├── db/            # Database connection and models
├── hash/          # Image hashing logic
├── routes/        # API routes (upload, search)
├── static/uploads # Uploaded image folder
├── main.go        # Entry point
└── go.mod         # Go module file



![Image](https://github.com/user-attachments/assets/8eb532cc-8c0a-4750-83f4-9d0f85d6742f)




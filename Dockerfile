# Sử dụng hình ảnh base chính thức của Golang
FROM golang:1.16-alpine as builder

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép go mod và sum files
COPY go.mod go.sum ./

# Tải về các phụ thuộc
RUN go mod download

# Sao chép mã nguồn vào container
COPY . .

# Biên dịch ứng dụng Go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Sử dụng scratch làm hình ảnh base thứ hai
FROM scratch

# Sao chép file thực thi từ builder
COPY --from=builder /app/main .

# Mở cổng 8080
EXPOSE 8080

# Chạy file thực thi
CMD ["./main"]
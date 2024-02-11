FROM golang:1.22

WORKDIR /app
COPY /grinha-de-backend-2024-q1-http3 /app/grinha-de-backend-2024-q1-http3
COPY /server-cert.pem /app/server-cert.pem
COPY /server.key /app/server.key

EXPOSE 8085/udp

EXPOSE 8085/tcp

CMD ["./grinha-de-backend-2024-q1-http3"]

import socket

ACK = 0x06
NACK = 0x15
STX = 0x02
ADDR = 0x80
ETX = 0x03

def start_tcp_server(host='localhost', port=65432):
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as server_socket:
        server_socket.bind((host, port))
        server_socket.listen()
        print(f"Server listening on {host}:{port}")
        
        while True:
            conn, addr = server_socket.accept()
            with conn:
                print(f"Connected by {addr}")
                while True:
                    try:
                        data = conn.recv(1024)
                        if not data:
                            break                        
                        response_buffer = bytearray()
                        response_buffer.append(STX)
                        response_buffer.append(ADDR)
                        response_buffer.append(ACK)
                        response_buffer.append(ETX)
                        response_buffer.append(0x38)
                        response_buffer.append(0x35)
                        print(response_buffer)
                        conn.sendall(response_buffer)
                    except:
                        break

if __name__ == "__main__":
    start_tcp_server()
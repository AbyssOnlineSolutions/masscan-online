# (1)コンパイラ
CC  = go build

TARGET  = masscan-online
# (4)コンパイル対象のソースコード
SRCS    = *.go

$(TARGET):
	$(CC) -o $(TARGET) $(SRCS)

all: clean  $(TARGET)

clean:
	-rm -f $(TARGET) 

run:
	./$(TARGET)
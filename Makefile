TGO = tinygo
OUT_FILE = out.hex
ENTRYPOINT = main.go
TGO_FLAGS = -target=pico

OCD = openocd
OCD_FLAGS = -f interface/picoprobe.cfg -f target/rp2040.cfg

all: flash

flash: $(OUT_FILE)
	$(OCD) $(OCD_FLAGS) -c "program $(OUT_FILE) reset verify exit"

$(OUT_FILE): $(ENTRYPOINT)
	$(TGO) build -o $(OUT_FILE) $(TGO_FLAGS) $(ENTRYPOINT)

clean:
	rm $(output)

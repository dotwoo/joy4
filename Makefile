.ONESHELL:
.PHONY: download build clean

# Temporary folder for building ffmpeg
TMP_DIR?=/tmp/

# Download AlexeyAB version of ffmpeg
download:
	rm -rf $(TMP_DIR)install_ffmpeg
	mkdir $(TMP_DIR)install_ffmpeg
	cd $(TMP_DIR)install_ffmpeg
	wget http://ffmpeg.org/releases/ffmpeg-4.4.tar.gz
	tar zxf ffmpeg-4.4.tar.gz
	cd -

# Build AlexeyAB version of ffmpeg for usage with CPU only.
build:
	cd $(TMP_DIR)install_ffmpeg/ffmpeg-4.4
	./configure --enable-shared \
    --enable-runtime-cpudetect \
    --enable-gpl \
    --enable-small \
    --enable-cross-compile \
    --disable-debug \
    --disable-static \
    --disable-doc \
    --disable-asm \
    --disable-ffmpeg \
    --disable-ffplay \
    --disable-ffprobe \
    --disable-postproc \
    --disable-avdevice \
    --disable-symver \
    --disable-stripping 
	make
	cd -



# Install system wide.
sudo_install:
	cd $(TMP_DIR)install_ffmpeg/ffmpeg-4.4
	make install
	cd -

# Cleanup temporary files for building process
clean:
	rm -rf $(TMP_DIR)install_ffmpeg

# Do every step for CPU-based only build.
install: download build sudo_install clean

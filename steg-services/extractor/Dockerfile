FROM ubuntu:24.04

ARG DEFAULT_WORKING_DIR="/analysis"

WORKDIR /tmp

# Update apt und installiere benötigte Pakete
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    wget \
    git \
    ca-certificates \
    libjpeg-dev \
    autoconf \
    automake \
    libtool \
    pkg-config \
    steghide \
    python3 \
    python3-pip \
    python3-venv \
    ffmpeg \
    libsm6 \
    libxext6 \
    binwalk \
    foremost \
    exiftool \
    p7zip-full \
    libffi-dev \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Erstelle das Arbeitsverzeichnis
RUN mkdir -p ${DEFAULT_WORKING_DIR} && chmod 777 ${DEFAULT_WORKING_DIR}
WORKDIR ${DEFAULT_WORKING_DIR}

# Setze die Umgebungsvariable für binwalk
ENV BINWALK_RM_EXTRACTION_SYMLINK=1

# Wechsle ins Arbeitsverzeichnis und kopiere die Anwendung
WORKDIR /app
COPY . .

# Erlaube pip systemweite Installationen
RUN mkdir -p $HOME/.config/pip && echo "[global]" > $HOME/.config/pip/pip.conf && echo "break-system-packages = true" >> $HOME/.config/pip/pip.conf

# Installiere Python-Abhängigkeiten
RUN pip install --no-cache-dir -r requirements.txt
RUN pip install --no-cache-dir -r LSBSteganography/requirements.txt

CMD ["python3", "-u", "server.py"]

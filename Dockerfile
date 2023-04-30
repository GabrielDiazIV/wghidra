FROM ubuntu
RUN apt-get update
RUN apt-get install -y openjdk-11-jdk
RUN apt-get install -y curl
RUN apt-get install -y unzip
COPY container /container
WORKDIR /container
RUN curl -L https://github.com/NationalSecurityAgency/ghidra/releases/download/Ghidra_10.2.3_build/ghidra_10.2.3_PUBLIC_20230208.zip > ghidra.zip
RUN unzip ghidra.zip
RUN rm ghidra.zip
RUN mv ghidra* ghidra
RUN apt-get install -y python3-pip
RUN pip install ghidra_bridge
CMD ["tail", "-f", "/dev/null"]
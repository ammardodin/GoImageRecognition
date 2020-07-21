FROM ctava/tfcgo
RUN mkdir -p /model && \
  curl -o /model/inception5h.zip -s "http://download.tensorflow.org/models/inception5h.zip" && \
  unzip /model/inception5h.zip -d /model
RUN go get -d "github.com/stretchr/testify/assert"

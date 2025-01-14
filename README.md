# Kafka Pipeline for VCS Hackathon 2024

This project was developed for the **VCS Hackathon 2024**. It is designed to:
- Consume data from Kafka.
- Transform messages into a suitable format.
- Send the transformed data to **Dify** for training purposes.

This pipeline showcases real-time data processing capabilities, combining flexible message transformation with a seamless integration to AI training systems like Dify.

---

## Features

- **Real-Time Data Consumption**: Processes data from Kafka topics.
- **Message Transformation**: Converts messages into a format compatible with Dify.
- **YAML Configuration**: Centralized configuration file for controlling pipeline behavior.

---

## Data Flow Diagram

Here is the data flow for the pipeline:

1. **Kafka**: Acts as the input source, producing messages.
2. **Pipeline**:
   - **Filter**: Validates incoming messages.
   - **Transform**: Converts message data to the required format.
3. **Dify**: Receives the transformed messages to train AI models.

```
Kafka Topic --> [Filter] --> [Transform] --> Dify Training API
```

---

## Configuration

The pipeline behavior is controlled through a YAML configuration file.

### Example `app.yml`:
```yaml
source:
  type: kafka
  kafka:
    address: kafka-server:9092
    topic: input-topic
    group: my-consumer-group

output:
  url: http://dify-training-system/api
```

### Key Sections:
- **`source`**: Defines the input source type and its specific configuration.
  - `type`: Input source type (`kafka`).
  - `kafka`: Kafka-specific configuration (address, topic, consumer group).
- **`output`**: Defines the API endpoint for Dify.

---

## Installation

1. **Clone the Repository**:

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```

---

## Usage

1. Configure your input and output in `app.yaml`.
2. Run the application using `go run cmd/forwarder`.
3. Check logs to monitor the pipeline's behavior.

---

## Extending the Pipeline

### Adding a New Input Source
1. Implement the `Source` interface in the `input/` package.
2. Add the new source type to the factory in `factory.go`.

### Adding a New Extractor
1. Implement the `Filter` interface in the `filter/` package.
2. Integrate the new extractor into the message processing flow in `main.go`.

### Adding a New Transformer
1. Implement the `Transformer` interface in the `transform/` package.
2. Replace or chain transformers in the message processing flow in `main.go`.

---



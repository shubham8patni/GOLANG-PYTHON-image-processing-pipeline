### Event-Driven Project: Real-time Image Processing Pipeline

**Components:** \

***
* Golang Microservices:\
_Image Uploader:_ \
    * Accepts user-uploaded images through an API endpoint (RESTful).
    * Validates and transforms images as needed (e.g., resize, convert format, storing, updating DB records).
    * Publishes a message containing image metadata (e.g., filename, size, format) to the image_upload topic in Kafka.
***


* Django Microservices:\
_Image Processor:_\
    * Subscribes to the image_upload topic in Kafka.
    * Consumes messages containing image metadata.
    * Fetches the image from a storage location (e.g., local file system, cloud storage) based on the metadata.
    * Applies desired processing (e.g., resize, apply filters, generate thumbnails).
    * Stores information in DB.
    * Publishes a new message containing the processed image data (e.g., byte array) to the image_processed topic in Kafka.
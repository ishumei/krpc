namespace go shumei.strategy.re

struct Image {
    1: optional binary content,
    2: optional i32 width,
    3: optional i32 height,
}

struct ImagePredictRequest {
    1: optional string requestId,
    2: optional string organization,
    3: optional Image image,
}

struct ImagePredictResult {
    1: optional string result,
}

service ImagePredictor {
    ImagePredictResult predict(1: ImagePredictRequest request),
    bool health(),
}
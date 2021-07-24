package request

type SatellitesRequest struct {
    Satellites []SateliteRequest `json:"satellites"`
}

type SateliteRequest struct {
    Name string `json:"name"`
    Distance float32 `json:"distance"`
    Message []string `json:"message"`
}
## kepler-demo

### List of APIs Available: 
 
   - "/metrics" - (GET) List of metrics available
        ```
        [{
            "id" : "1",
            "name" : "ABC"
        },
        {
            "id" : "2",
            "name" : "DEF"
        }]
        ```
    
   - "/metrics/{id}/recommendations" - (GET) List of metrics as per recommendation
        ```
        [{
            "id" : "1",
            "metric_id" : "1",
            "name" : "ABC"
        },
        {
            "id" : "6",
            "metric_id" : "1",
            "name" : "DEF"
        }]
        ```
   - "/metrics/{id}/data" - (GET) Data for the metrics
        ```
	{
    "type": "FeatureCollection",
    "features": [
        {
            "type": "Feature",
            "properties": {
                "id": 0.6046602879796196
            },
            "geometry": {
                "type": "Point",
                "coordinates": [
                    106.85,
                    106.85
                ]
            }
        },
        {
            "type": "Feature",
            "properties": {
                "id": 0.9405090880450124
            },
            "geometry": {
                "type": "Point",
                "coordinates": [
                    106.85,
                    106.85
                ]
            }
        }]
}	        
        ```
   - "/service-types" - (GET) List of service types available
        ```
        [{
          "id": 1,
          "name": "ABC"
        },
        {
          "id": 2,
          "name": "DEF"
        }]
        ```
   - "/service-areas" - (GET) List of service areas available 
        ```
        [{
          "id": 1,
          "name": "ABC"
        },
        {
          "id": 2,
          "name": "DEF"
        }]
        ```


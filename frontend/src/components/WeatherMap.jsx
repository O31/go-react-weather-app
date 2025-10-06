import { MapContainer, TileLayer, Marker, Popup, useMap, useMapEvents } from "react-leaflet"
import "leaflet/dist/leaflet.css"
import "./WeatherMap.css"
import { useEffect } from "react"

// Fix for default markers in react-leaflet
import L from "leaflet"
import icon from "leaflet/dist/images/marker-icon.png"
import iconShadow from "leaflet/dist/images/marker-shadow.png"

let DefaultIcon = L.divIcon({
  html: `<div class="custom-marker">📍</div>`,
  className: "custom-div-icon",
  iconSize: [25, 25],
  iconAnchor: [12, 25],
})

L.Marker.prototype.options.icon = DefaultIcon

function CenterMapView({ coords }) {
  const map = useMap()
  useEffect(() => {
    if (coords) {
      map.setView([coords.lat, coords.lng], map.getZoom(), { animate: true })
    }
  }, [coords, map])
  return null
}

function LocationMarker({ onLocationSelect, selectedLocation }) {
  useMapEvents({
    click(e) {
      const { lat, lng } = e.latlng
      onLocationSelect(lat, lng)
    },
  })

  return selectedLocation ? (
    <Marker position={[selectedLocation.lat, selectedLocation.lng]}>
      <Popup>
        Selected location: {selectedLocation.lat.toFixed(4)}, {selectedLocation.lng.toFixed(4)}
      </Popup>
    </Marker>
  ) : null
}

function WeatherMap({ onLocationSelect, selectedLocation }) {
  return (
    <div className="map-container">
      <h4>Click on the map to get weather for that location</h4>
      <MapContainer
        center={[selectedLocation.lat, selectedLocation.lng]}
        zoom={6}
        className="leaflet-container"
        scrollWheelZoom={true}
        style={{ height: "300px", minHeight: "300px", width: "100%" }}
      >
        <TileLayer
          attribution="&copy; OpenStreetMap contributors"
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        {/* Center the map when selectedLocation changes */}
        <CenterMapView coords={selectedLocation} />

        <LocationMarker onLocationSelect={onLocationSelect} selectedLocation={selectedLocation} />
      </MapContainer>
    </div>
  )
}

export default WeatherMap

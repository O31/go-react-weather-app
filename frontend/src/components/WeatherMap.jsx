import React from "react"
import { MapContainer, TileLayer, useMapEvents, Marker, Popup } from "react-leaflet"
import "leaflet/dist/leaflet.css"
import "./WeatherMap.css"

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
        center={[59.3293, 18.0686]} // Stockholm coordinates
        zoom={6}
        className="leaflet-container"
      >
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        <LocationMarker onLocationSelect={onLocationSelect} selectedLocation={selectedLocation} />
      </MapContainer>
    </div>
  )
}

export default WeatherMap

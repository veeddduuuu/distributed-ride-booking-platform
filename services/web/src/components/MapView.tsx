import { MapContainer, TileLayer, Marker, Polyline, useMapEvents } from 'react-leaflet';
import polyline from '@mapbox/polyline';
import L from 'leaflet';
import markerIcon2x from 'leaflet/dist/images/marker-icon-2x.png';
import markerIcon from 'leaflet/dist/images/marker-icon.png';
import markerShadow from 'leaflet/dist/images/marker-shadow.png';
import type { Coordinates, Route } from '../api/trip';

// Leaflet's default marker icons reference relative asset paths that break when
// bundled by Vite, leaving markers invisible. Point them at the bundled URLs.
L.Icon.Default.mergeOptions({
  iconRetinaUrl: markerIcon2x,
  iconUrl: markerIcon,
  shadowUrl: markerShadow,
});

interface MapViewProps {
  pickup: Coordinates | null;
  destination: Coordinates | null;
  route: Route | null;
  onMapClick: (lat: number, lng: number) => void;
}

function ClickHandler({ onClick }: { onClick: (lat: number, lng: number) => void }) {
  useMapEvents({
    click(e) {
      onClick(e.latlng.lat, e.latlng.lng);
    },
  });
  return null;
}

export default function MapView({ pickup, destination, route, onMapClick }: MapViewProps) {
  // Decode the OSRM polyline string into an array of [lat, lng] pairs
  const decodedPath = route?.polyline 
    ? polyline.decode(route.polyline) 
    : [];

  return (
    <MapContainer center={[18.5204, 73.8567]} zoom={11} style={{ height: '100%', width: '100%' }}>
      <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
      <ClickHandler onClick={onMapClick} />
      
      {pickup && <Marker position={[pickup.latitude, pickup.longitude]} />}
      {destination && <Marker position={[destination.latitude, destination.longitude]} />}
      
      {decodedPath.length > 0 && <Polyline positions={decodedPath} color="blue" weight={5} />}
    </MapContainer>
  );
}   
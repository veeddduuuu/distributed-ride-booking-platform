import { MapContainer, TileLayer, Marker, Polyline, useMapEvents } from 'react-leaflet';
import polyline from '@mapbox/polyline';
import L from 'leaflet';
import type { Coordinates, Route } from '../api/trip';
import FitBounds from './FitBounds';

// Create minimal SVG markers to replace default Leaflet pins
const createSvgMarker = (color: string, filled: boolean) => {
  const svg = `
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20">
      <circle cx="12" cy="12" r="8" fill="${filled ? color : 'white'}" stroke="${color}" stroke-width="4" />
    </svg>
  `;
  return L.divIcon({
    html: svg,
    className: 'custom-svg-marker',
    iconSize: [20, 20],
    iconAnchor: [10, 10],
  });
};

const pickupIcon = createSvgMarker('#0D9488', true); // Filled teal
const destIcon = createSvgMarker('#0D9488', false); // Outlined teal

interface MapViewProps {
  pickup: Coordinates | null;
  destination: Coordinates | null;
  route: Route | null;
  onMapClick: (lat: number, lng: number) => void;
}

function ClickHandler({ onClick }: { onClick: (lat: number, lng: number) => void }) {
  const map = useMapEvents({
    click(e) {
      onClick(e.latlng.lat, e.latlng.lng);
    },
    mousemove() {
      // Small UX detail: cursor becomes pointer only when map is clickable (i.e. not panning)
      map.getContainer().style.cursor = 'crosshair';
    }
  });
  return null;
}

export default function MapView({ pickup, destination, route, onMapClick }: MapViewProps) {
  // Decode the OSRM polyline string into an array of [lat, lng] pairs
  const decodedPath = route?.polyline 
    ? polyline.decode(route.polyline) 
    : [];

  return (
    <MapContainer 
      center={[18.5204, 73.8567]} 
      zoom={11} 
      style={{ height: '100%', width: '100%' }}
      zoomControl={false}
    >
      {/* CartoDB Voyager tiles - cleaner, minimal, fits the design system well */}
      <TileLayer url="https://{s}.basemaps.cartocdn.com/rastertiles/voyager/{z}/{x}/{y}{r}.png" />
      <ClickHandler onClick={onMapClick} />
      
      {pickup && <Marker position={[pickup.latitude, pickup.longitude]} icon={pickupIcon} />}
      {destination && <Marker position={[destination.latitude, destination.longitude]} icon={destIcon} />}
      
      {decodedPath.length > 0 && (
        <>
          <Polyline 
            positions={decodedPath} 
            color="#0D9488" 
            weight={4}
            opacity={0.8}
            lineCap="round"
            lineJoin="round"
          />
          <FitBounds positions={decodedPath as [number, number][]} />
        </>
      )}
    </MapContainer>
  );
}
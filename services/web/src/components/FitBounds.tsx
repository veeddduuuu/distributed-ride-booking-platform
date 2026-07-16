import { useEffect } from 'react';
import { useMap } from 'react-leaflet';
import L from 'leaflet';

interface FitBoundsProps {
  positions: [number, number][];
}

/**
 * Adjusts the map viewport to fit all positions with comfortable padding.
 * Drop this inside a <MapContainer> alongside the <Polyline>.
 */
export default function FitBounds({ positions }: FitBoundsProps) {
  const map = useMap();

  useEffect(() => {
    if (positions.length < 2) return;

    const bounds = L.latLngBounds(positions);
    map.fitBounds(bounds, { padding: [50, 50], maxZoom: 15 });
  }, [map, positions]);

  return null;
}

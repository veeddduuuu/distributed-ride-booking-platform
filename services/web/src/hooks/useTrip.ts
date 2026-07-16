import { useState, useEffect, useCallback } from 'react';
import type {
  Coordinates,
  Route,
  RideShare,
  TripPreviewResponse,
} from '../api/trip';
import { fetchTripPreview } from '../api/trip';

export function useTrip() {
  const [pickup, setPickup] = useState<Coordinates | null>(null);
  const [destination, setDestination] = useState<Coordinates | null>(null);
  const [tripPreview, setTripPreview] = useState<TripPreviewResponse | null>(
    null
  );
  const [selectedFare, setSelectedFare] = useState<RideShare | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSetPickup = useCallback((coords: Coordinates) => {
    setPickup((prevPickup) => {
      if (!prevPickup) return coords;
      return prevPickup;
    });
  }, []);

  const handleMapClick = useCallback(
    (coords: Coordinates) => {
      if (!pickup) {
        setPickup(coords);
      } else if (!destination) {
        setDestination(coords);
      } else {
        // Reset: new pickup
        setPickup(coords);
        setDestination(null);
        setTripPreview(null);
        setSelectedFare(null);
        setError(null);
      }
    },
    [pickup, destination]
  );

  const selectFare = useCallback((fare: RideShare) => {
    setSelectedFare(fare);
  }, []);

  const resetTrip = useCallback(() => {
    setPickup(null);
    setDestination(null);
    setTripPreview(null);
    setSelectedFare(null);
    setIsLoading(false);
    setError(null);
  }, []);

  // Fetch trip preview when both points are set
  useEffect(() => {
    if (!pickup || !destination) return;

    let cancelled = false;
    setIsLoading(true);
    setError(null);
    setSelectedFare(null);

    fetchTripPreview(pickup, destination)
      .then((preview) => {
        if (!cancelled) {
          setTripPreview(preview);
        }
      })
      .catch((err) => {
        if (!cancelled) {
          console.error('Error fetching trip preview:', err);
          setError(err instanceof Error ? err.message : 'Failed to load route');
        }
      })
      .finally(() => {
        if (!cancelled) {
          setIsLoading(false);
        }
      });

    return () => {
      cancelled = true;
    };
  }, [pickup, destination]);

  // Derive route for backward compat
  const route: Route | null = tripPreview?.route ?? null;

  return {
    pickup,
    destination,
    route,
    tripPreview,
    selectedFare,
    isLoading,
    error,
    handleMapClick,
    handleSetPickup,
    selectFare,
    resetTrip,
  };
}

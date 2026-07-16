export interface Coordinates {
  latitude: number;
  longitude: number;
}

export interface Route {
  distance: number; // in meters
  duration: number; // in seconds
  polyline: string;
}

export interface RideShare {
  id: string;
  userId: string;
  packageSlug: string;
  totalPrice: number;
}

export interface TripPreviewResponse {
  tripId: string;
  route: Route;
  rideFares: RideShare[];
}

export const fetchTripPreview = async (
  pickup: Coordinates,
  destination: Coordinates
): Promise<TripPreviewResponse> => {
  const response = await fetch('/trip/preview', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      userId: '123',
      pickup,
      destination,
    }),
  });

  if (!response.ok) {
    throw new Error(`Error fetching trip preview: ${response.statusText}`);
  }

  const data: TripPreviewResponse = await response.json();
  return data;
};
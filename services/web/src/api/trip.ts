export interface Coordinates {
    latitude: number;
    longitude: number;
}

export interface Route {
    distance: number; // in meters
    duration: number; // in seconds
    polyline : string;
}

export const fetchTripPreview = async (pickup: Coordinates, destination: Coordinates) : Promise<Route> => {
    // Relative path: served same-origin in production behind the gateway, and
    // proxied to the gateway in dev via vite.config.ts. `api-gateway` is a
    // Docker-internal hostname the browser cannot resolve.
    const response = await fetch('/trip/preview', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            userId: '123',
            pickup,
            destination
        })
    }); 

    if (!response.ok) {
        throw new Error(`Error fetching trip preview: ${response.statusText}`);
    }
    const data : Route = await response.json();
    return data;
}
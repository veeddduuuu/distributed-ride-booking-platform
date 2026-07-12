import { useState, useEffect, useRef } from 'react';

const WS_URL = 'ws://localhost:8080/ws/driver';

export interface DriverData {
  userId: string;
  name: string;
  carNumber: string;
  profilePicture: string;
  packageSlug: string;
}

export function useDriverWebSocket(userId: string, packageSlug: string) {
  const [isConnected, setIsConnected] = useState(false);
  const [driverInfo, setDriverInfo] = useState<DriverData | null>(null);
  const [availableRides, setAvailableRides] = useState<any[]>([]);
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!userId || !packageSlug) return;

    const connect = () => {
      ws.current = new WebSocket(`${WS_URL}?userId=${userId}&packageSlug=${packageSlug}`);

      ws.current.onopen = () => {
        console.log('Driver WS Connected');
        setIsConnected(true);
      };

      ws.current.onclose = () => {
        console.log('Driver WS Disconnected');
        setIsConnected(false);
        setDriverInfo(null);
      };

      ws.current.onerror = (error) => {
        console.error('Driver WS Error:', error);
      };

      ws.current.onmessage = (event) => {
        try {
          const message = JSON.parse(event.data);
          console.log('Driver WS Message Received:', message);

          if (message.type === 'driver_registered') {
            setDriverInfo(message.data);
          } else if (message.type === 'new_ride') {
            setAvailableRides(prev => [...prev, message.data]);
          }
        } catch (e) {
          console.error('Failed to parse WS message:', e);
        }
      };
    };

    connect();

    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, [userId, packageSlug]);

  return { isConnected, driverInfo, availableRides };
}

import { useState, useEffect, useCallback, useRef } from 'react';

const WS_URL = 'ws://localhost:8080/ws/rider';

export function useRiderWebSocket(userId: string) {
  const [isConnected, setIsConnected] = useState(false);
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    if (!userId) return;

    const connect = () => {
      ws.current = new WebSocket(`${WS_URL}?userId=${userId}`);

      ws.current.onopen = () => {
        console.log('Rider WS Connected');
        setIsConnected(true);
      };

      ws.current.onclose = () => {
        console.log('Rider WS Disconnected');
        setIsConnected(false);
      };

      ws.current.onerror = (error) => {
        console.error('Rider WS Error:', error);
      };

      ws.current.onmessage = (event) => {
        console.log('Rider WS Message Received:', event.data);
      };
    };

    connect();

    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, [userId]);

  const sendRideRequest = useCallback((data: any) => {
    if (ws.current && isConnected) {
      ws.current.send(JSON.stringify({ type: 'ride_request', data }));
    } else {
      console.warn('Rider WS not connected. Cannot send request.');
    }
  }, [isConnected]);

  return { isConnected, sendRideRequest };
}

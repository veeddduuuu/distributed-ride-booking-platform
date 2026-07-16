interface StatusBadgeProps {
  connected: boolean;
}

export default function StatusBadge({ connected }: StatusBadgeProps) {
  return (
    <span className="status-badge">
      <span
        className={`status-dot ${
          connected ? 'status-dot--connected' : 'status-dot--disconnected'
        }`}
      />
      {connected ? 'Connected' : 'Disconnected'}
    </span>
  );
}

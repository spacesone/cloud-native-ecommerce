import { useState } from 'react';
import { toast } from 'sonner';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { orders } from '@/api/api';

export default function OrderManagement() {
  const [orderId, setOrderId] = useState('');
  const [loading, setLoading] = useState(false);

  const handleCancel = async () => {
    if (!orderId) {
      toast.error('Please enter an order ID');
      return;
    }
    
    setLoading(true);
    try {
      await orders.cancel(orderId);
      toast.success('Order cancelled successfully');
      setOrderId('');
    } catch (error) {
      toast.error('Failed to cancel order');
      console.error('Error cancelling order:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleComplete = async () => {
    if (!orderId) {
      toast.error('Please enter an order ID');
      return;
    }

    setLoading(true);
    try {
      await orders.complete(orderId);
      toast.success('Order completed successfully');
      setOrderId('');
    } catch (error) {
      toast.error('Failed to complete order');
      console.error('Error completing order:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleCheckStock = async () => {
    if (!orderId) {
      toast.error('Please enter an order ID');
      return;
    }

    setLoading(true);
    try {
      const response = await orders.getStock(orderId);
      toast.success('Stock retrieved successfully');
      console.log('Stock information:', response.data);
    } catch (error) {
      toast.error('Failed to retrieve stock information');
      console.error('Error getting stock:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto py-6 space-y-4">
      <Card>
        <CardHeader>
          <CardTitle>Order Management</CardTitle>
          <CardDescription>
            Manage order statuses and check stock information
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="flex items-center space-x-4">
            <Input
              placeholder="Enter Order ID"
              value={orderId}
              onChange={(e) => setOrderId(e.target.value)}
              disabled={loading}
            />
          </div>
          <div className="flex space-x-4">
            <Button
              variant="destructive"
              onClick={handleCancel}
              disabled={loading || !orderId}
            >
              Cancel Order
            </Button>
            <Button
              variant="default"
              onClick={handleComplete}
              disabled={loading || !orderId}
            >
              Complete Order
            </Button>
            <Button
              variant="outline"
              onClick={handleCheckStock}
              disabled={loading || !orderId}
            >
              Check Stock
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

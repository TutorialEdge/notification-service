
ALTER TABLE subscribers 
ADD COLUMN created_at timestamp DEFAULT CURRENT_DATE,
ADD COLUMN updated_at timestamp DEFAULT CURRENT_DATE,
ADD COLUMN deleted_at timestamp;

ALTER TABLE notifications 
ADD COLUMN created_at timestamp DEFAULT CURRENT_DATE,
ADD COLUMN updated_at timestamp DEFAULT CURRENT_DATE,
ADD COLUMN deleted_at timestamp;

ALTER TABLE list 
ADD COLUMN created_at timestamp DEFAULT CURRENT_DATE,
ADD COLUMN updated_at timestamp DEFAULT CURRENT_DATE,
ADD COLUMN deleted_at timestamp;
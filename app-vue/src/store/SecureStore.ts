import SecureLS from 'secure-ls';

export class SecureStorage implements Storage {
  private ls: SecureLS;

  constructor() {
    this.ls = new SecureLS({ encodingType: 'aes' });
  }

  get length(): number {
    return this.ls.getAllKeys().length;
  }

  key(index: number): string | null {
    const keys = this.ls.getAllKeys();
    return keys[index] || null;
  }

  getItem(key: string): string | null {
    return this.ls.get(key) || null;
  }

  setItem(key: string, value: string): void {
    this.ls.set(key, value);
  }

  removeItem(key: string): void {
    this.ls.remove(key);
  }

  clear(): void {
    this.ls.removeAll();
  }
}

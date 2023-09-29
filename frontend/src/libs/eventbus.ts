type TypeMap = { [key: string]: Array<(payload?: any[]) => void> };

export class EventEmitter {
  events: TypeMap;
  constructor() {
    this.events = {};
  }
  on(event: string, fn: (...payload: any[]) => void) {
    if (!this.events[event]) this.events[event] = [];
    this.events[event].push(fn);
  }
  emit(event: string, ...payload: any[]) {
    this.events[event].forEach((fn) => fn(...payload));
  }
}

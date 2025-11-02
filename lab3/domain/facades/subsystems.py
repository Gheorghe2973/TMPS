from domain.models.computer import Computer


class ComponentValidator:
    def validate_cpu_ram_compatibility(self, cpu: str, ram: str) -> bool:
        ddr5_cpus = ["Intel i9", "AMD Ryzen 9", "Intel i7 13th"]
        if any(cpu_name in cpu for cpu_name in ddr5_cpus):
            return "DDR5" in ram
        return "DDR4" in ram
    
    def validate_gpu_power(self, gpu: str, ram: str) -> bool:
        high_end_gpus = ["RTX 4090", "RTX 4080", "RTX 4070"]
        if any(gpu_name in gpu for gpu_name in high_end_gpus):
            ram_size = int(''.join(filter(str.isdigit, ram)))
            return ram_size >= 32
        return True


class PriceCalculator:
    def __init__(self):
        self._prices = {
            "Intel i9": 500, "Intel i7": 350, "Intel i5": 200,
            "AMD Ryzen 9": 450, "AMD EPYC": 800,
            "RTX 4090": 1600, "RTX 4080": 1200, "RTX 4070": 600,
            "32GB DDR5": 200, "16GB DDR4": 80, "64GB DDR5": 400,
            "128GB ECC": 800, "2TB NVMe SSD": 200, "1TB SSD": 100,
            "512GB SSD": 60, "4TB RAID": 600
        }
    
    def calculate_component_price(self, component: str) -> float:
        for key, price in self._prices.items():
            if key in component:
                return price
        return 50.0
    
    def calculate_total_price(self, computer: Computer) -> float:
        total = 0.0
        if computer.cpu:
            total += self.calculate_component_price(computer.cpu)
        if computer.ram:
            total += self.calculate_component_price(computer.ram)
        if computer.storage:
            total += self.calculate_component_price(computer.storage)
        if computer.gpu:
            total += self.calculate_component_price(computer.gpu)
        if computer.case:
            total += 100.0
        total += len(computer.accessories) * 50.0
        return total


class AssemblyScheduler:
    def __init__(self):
        self._queue = []
    
    def schedule_assembly(self, order_id: str, computer: Computer) -> str:
        position = len(self._queue) + 1
        self._queue.append(order_id)
        
        complexity = sum([
            1 if computer.cpu else 0,
            1 if computer.ram else 0,
            1 if computer.storage else 0,
            1 if computer.gpu else 0,
            1 if computer.case else 0,
            len(computer.accessories)
        ])
        
        hours = complexity * 0.5 + 1
        return f"Order #{order_id} scheduled. Position: {position}. Estimated: {hours}h"
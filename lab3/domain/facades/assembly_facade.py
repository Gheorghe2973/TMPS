from typing import Dict
from domain.models.computer import Computer
from domain.facades.subsystems import ComponentValidator, PriceCalculator, AssemblyScheduler


class ComputerAssemblyFacade:
    def __init__(self):
        self._validator = ComponentValidator()
        self._price_calculator = PriceCalculator()
        self._scheduler = AssemblyScheduler()
        self._order_counter = 1000
    
    def order_computer(self, computer: Computer) -> Dict[str, any]:
        order_id = str(self._order_counter)
        self._order_counter += 1
        
        is_valid = True
        validation_messages = []
        
        if computer.cpu and computer.ram:
            if not self._validator.validate_cpu_ram_compatibility(computer.cpu, computer.ram):
                is_valid = False
                validation_messages.append("⚠ CPU and RAM may not be compatible")
        
        if computer.gpu and computer.ram:
            if not self._validator.validate_gpu_power(computer.gpu, computer.ram):
                is_valid = False
                validation_messages.append("⚠ Insufficient RAM for this GPU")
        
        total_price = self._price_calculator.calculate_total_price(computer)
        schedule_info = self._scheduler.schedule_assembly(order_id, computer)
        
        return {
            "order_id": order_id,
            "computer": str(computer),
            "is_valid": is_valid,
            "validation_messages": validation_messages,
            "total_price": f"${total_price:.2f}",
            "schedule": schedule_info,
            "status": "✓ Order placed successfully" if is_valid else "⚠ Order placed with warnings"
        }
    
    def get_quick_quote(self, computer: Computer) -> str:
        price = self._price_calculator.calculate_total_price(computer)
        return f"${price:.2f}"
from abc import ABC, abstractmethod
from domain.models.computer import Computer


class ComputerComponent(ABC):
    @abstractmethod
    def get_description(self) -> str:
        pass
    
    @abstractmethod
    def get_cost(self) -> float:
        pass


class BaseComputer(ComputerComponent):
    def __init__(self, computer: Computer):
        self._computer = computer
        self._base_cost = 1000.0
    
    def get_description(self) -> str:
        return str(self._computer)
    
    def get_cost(self) -> float:
        return self._base_cost


class ComputerDecorator(ComputerComponent):
    def __init__(self, computer_component: ComputerComponent):
        self._component = computer_component
    
    @abstractmethod
    def get_description(self) -> str:
        pass
    
    @abstractmethod
    def get_cost(self) -> float:
        pass


class WarrantyDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent, years: int = 3):
        super().__init__(computer_component)
        self._years = years
        self._warranty_cost = years * 50.0
    
    def get_description(self) -> str:
        return f"{self._component.get_description()} + {self._years}Y Warranty"
    
    def get_cost(self) -> float:
        return self._component.get_cost() + self._warranty_cost


class OverclockingDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent):
        super().__init__(computer_component)
        self._overclock_cost = 200.0
    
    def get_description(self) -> str:
        return f"{self._component.get_description()} + Professional Overclocking"
    
    def get_cost(self) -> float:
        return self._component.get_cost() + self._overclock_cost


class RGBLightingDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent):
        super().__init__(computer_component)
        self._rgb_cost = 150.0
    
    def get_description(self) -> str:
        return f"{self._component.get_description()} + RGB Lighting System"
    
    def get_cost(self) -> float:
        return self._component.get_cost() + self._rgb_cost


class CustomPaintDecorator(ComputerDecorator):
    def __init__(self, computer_component: ComputerComponent, color: str):
        super().__init__(computer_component)
        self._color = color
        self._paint_cost = 300.0
    
    def get_description(self) -> str:
        return f"{self._component.get_description()} + Custom {self._color} Paint"
    
    def get_cost(self) -> float:
        return self._component.get_cost() + self._paint_cost
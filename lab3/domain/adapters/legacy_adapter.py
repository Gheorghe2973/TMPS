from abc import ABC, abstractmethod


class LegacyComponent:
    def __init__(self, component_type: str, specifications: str):
        self.component_type = component_type
        self.specifications = specifications
    
    def get_component_info(self) -> str:
        return f"{self.component_type}: {self.specifications}"


class ModernComponentInterface(ABC):
    @abstractmethod
    def get_name(self) -> str:
        pass
    
    @abstractmethod
    def get_specs(self) -> str:
        pass
    
    @abstractmethod
    def is_compatible(self) -> bool:
        pass


class LegacyComponentAdapter(ModernComponentInterface):
    def __init__(self, legacy_component: LegacyComponent):
        self._legacy_component = legacy_component
        self._compatibility_check()
    
    def get_name(self) -> str:
        return self._legacy_component.component_type
    
    def get_specs(self) -> str:
        return self._legacy_component.specifications
    
    def is_compatible(self) -> bool:
        return self._compatible
    
    def _compatibility_check(self):
        self._compatible = bool(self._legacy_component.specifications)
    
    def __str__(self):
        status = "✓ Compatible" if self._compatible else "✗ Incompatible"
        return f"{self.get_name()}: {self.get_specs()} [{status}]"